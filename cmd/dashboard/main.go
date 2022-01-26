package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/brumhard/alligotor"
	"github.com/brumhard/pr-dashboard/app"
	"github.com/brumhard/pr-dashboard/pkg/api"
	ownhttp "github.com/brumhard/pr-dashboard/pkg/http"
	dashboardv1 "github.com/brumhard/pr-dashboard/pkg/pb/dashboard/v1"
	"github.com/brumhard/pr-dashboard/pkg/pr"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// This controls the maxprocs environment variable in container runtimes.
	// see https://martin.baillie.id/wrote/gotchas-in-the-go-network-packages-defaults/#bonus-gomaxprocs-containers-and-the-cfs
	_ "go.uber.org/automaxprocs"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}

type Config struct {
	GRPCAddr        string
	GRPCWebAddr     string
	LogLevel        zapcore.Level
	Providers       []pr.ProviderConfiguration
	RefreshInterval time.Duration
}

func run() error {
	cfgLoader := alligotor.New(
		alligotor.NewFilesSource("./configs/config.yaml", "/Users/brumhard/repos/personal/pr-dashboard/configs/config.yaml"),
		alligotor.NewEnvSource("helpr"),
	)
	config := Config{
		LogLevel:        zapcore.InfoLevel,
		GRPCWebAddr:     ":8080",
		GRPCAddr:        ":8081",
		RefreshInterval: 2 * time.Minute,
	}
	if err := cfgLoader.Get(&config); err != nil {
		return err
	}

	loggerCfg := zap.NewProductionConfig()
	loggerCfg.Level = zap.NewAtomicLevelAt(config.LogLevel)
	logger, err := loggerCfg.Build()
	if err != nil {
		return err
	}

	defer logger.Sync()

	aggregator, err := pr.NewAggregatorService(config.Providers)
	if err != nil {
		return err
	}

	cached, err := pr.WithAutoRefresh(logger.Named("autorefresh service"), aggregator, config.RefreshInterval)
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	dashboardv1.RegisterDashboardServiceServer(srv, api.NewGRPC(cached))
	reflection.Register(srv)
	wrapped := grpcweb.WrapServer(
		srv,
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)

	webContents, err := fs.Sub(app.Static, "build/web")
	if err != nil {
		return err
	}

	httpServer := &http.Server{
		// These interfere with websocket streams, disable for now
		// ReadTimeout: 5 * time.Second,
		// WriteTimeout: 10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
		Addr:              config.GRPCWebAddr,
		TLSConfig: &tls.Config{
			PreferServerCipherSuites: true,
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519,
			},
			MinVersion: tls.VersionTLS12,
		},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if wrapped.IsGrpcWebRequest(r) || wrapped.IsAcceptableGrpcCorsRequest(r) || wrapped.IsGrpcWebSocketRequest(r) {
				wrapped.ServeHTTP(w, r)
				return
			}
			ownhttp.NewSPAHandler(http.FS(webContents), "index.html").ServeHTTP(w, r)
		}),
	}

	listen, err := net.Listen("tcp", config.GRPCAddr)
	if err != nil {
		return err
	}

	ctx, ccancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer ccancel()

	eg, gctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		logger.Info("running grpc server", zap.String("addr", config.GRPCAddr))

		go srv.Serve(listen)
		<-gctx.Done()
		srv.GracefulStop()

		logger.Info("finished shutting down grpc server")
		return nil
	})

	eg.Go(func() error {
		logger.Info("running http server", zap.String("addr", config.GRPCWebAddr))

		go httpServer.ListenAndServe()
		<-gctx.Done()

		httpShutdownTimeout, scancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer scancel()
		err := httpServer.Shutdown(httpShutdownTimeout)

		logger.Info("finished shutting down http server", zap.Error(err))
		return err
	})

	return eg.Wait()
}
