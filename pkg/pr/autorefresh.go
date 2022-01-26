package pr

import (
	"context"
	"time"

	"go.uber.org/zap"
)

var _ StreamerService = (*AutoRefreshingService)(nil)

type AutoRefreshingService struct {
	logger          *zap.Logger
	service         Service
	currentPRs      []Repository
	joinc           chan chan []Repository
	leavec          chan chan []Repository
	prc             chan []Repository
	clients         map[chan []Repository]struct{}
	refreshInterval time.Duration
}

func WithAutoRefresh(logger *zap.Logger, service Service, refreshInterval time.Duration) (Service, error) {
	s := &AutoRefreshingService{
		logger:          logger,
		service:         service,
		joinc:           make(chan chan []Repository),
		leavec:          make(chan chan []Repository),
		prc:             make(chan []Repository),
		clients:         make(map[chan []Repository]struct{}),
		refreshInterval: refreshInterval,
	}
	s.run(context.TODO())
	return s, nil
}

func (s AutoRefreshingService) run(ctx context.Context) {
	go func() {
		for {
			go func() {
				start := time.Now()
				prs, err := s.service.GetAllPRs(ctx)
				if err != nil {
					s.logger.Error("failed fetching prs", zap.Error(err))
					return
				}
				s.logger.Info("refreshed prs", zap.Duration("duration", time.Since(start)))
				s.prc <- prs
			}()
			select {
			case <-ctx.Done():
				return
			case <-time.After(s.refreshInterval):
			}
		}
	}()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case clientc := <-s.joinc:
				s.logger.Debug("new client joining, checking whether cached result is available", zap.Bool("cacheReady", s.currentPRs != nil))
				if s.currentPRs != nil {
					go s.tryWriteToClient(ctx, clientc, s.currentPRs)
				}
				s.clients[clientc] = struct{}{}
			case clientc := <-s.leavec:
				s.logger.Debug("client leaving")
				delete(s.clients, clientc)
			case repos := <-s.prc:
				s.currentPRs = repos
				for clientc := range s.clients {
					go s.tryWriteToClient(ctx, clientc, repos)
				}
			}
		}
	}()
}

func (s AutoRefreshingService) tryWriteToClient(ctx context.Context, clientc chan []Repository, repos []Repository) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	select {
	case <-timeoutCtx.Done():
		s.logger.Warn("timed out while trying to send refresh to client")
	case clientc <- repos:
	}
}

func (s AutoRefreshingService) GetAllPRs(ctx context.Context) ([]Repository, error) {
	s.logger.Debug("Invoked GetAllPRs")
	clientc := make(chan []Repository)
	defer close(clientc)
	s.joinc <- clientc
	defer func() { s.leavec <- clientc }()

	repos := <-clientc
	return repos, nil
}

// StreamAllPRs returns a channel which will receive a new list of repositories every time
// the list is updated. It will run until the passed in context is closed.
func (s AutoRefreshingService) StreamAllPRs(ctx context.Context) (chan []Repository, error) {
	s.logger.Debug("Invoked StreamAllPRs")
	clientc := make(chan []Repository)
	s.joinc <- clientc

	go func() {
		defer close(clientc)
		defer func() { s.leavec <- clientc }()
		<-ctx.Done()
	}()

	return clientc, nil
}
