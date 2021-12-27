package main

import (
	"context"
	"fmt"
	"os"

	"github.com/brumhard/alligotor"
	"github.com/brumhard/pr-dashboard/pkg/pr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

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
	LogLevel  zapcore.Level
	Providers []pr.ProviderConfiguration
}

func run() error {
	cfgLoader := alligotor.New(alligotor.NewFilesSource("./configs/config.yaml"))
	var config Config
	if err := cfgLoader.Get(&config); err != nil {
		return err
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	defer logger.Sync()

	aggregator, err := pr.NewAggregator(config.Providers)
	if err != nil {
		return err
	}

	prs, err := aggregator.GetAllPRs(context.Background())
	if err != nil {
		return err
	}

	for _, pr := range prs {
		fmt.Println(pr.Title)
	}

	return nil
}
