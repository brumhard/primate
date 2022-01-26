package pr

import (
	"context"
	"encoding/json"
	"time"

	"github.com/allegro/bigcache/v3"
)

var _ StreamerService = (*AutoRefreshingService)(nil)

type AutoRefreshingService struct {
	service    Service
	currentPRs []Repository
	joinc      chan chan []Repository
	leavec     chan chan []Repository
	prc        chan []Repository
	clients    map[chan []Repository]struct{}
}

func WithAutoRefresh(service Service) (Service, error) {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(time.Minute))
	if err != nil {
		return nil, err
	}

	return &CachedService{service: service, cache: cache}, nil
}

func (s AutoRefreshingService) run(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case clientc := <-s.joinc:
				go tryWriteToClient(ctx, clientc, s.currentPRs)
				s.clients[clientc] = struct{}{}
			case clientc := <-s.leavec:
				delete(s.clients, clientc)
			case repos := <-s.prc:
				if s.currentPRs == repos {
					continue
				}
				s.currentPRs = repos
				for clientc := range s.clients {
					go tryWriteToClient(ctx, clientc, repos)
				}
			}
		}
	}()
}

func tryWriteToClient(ctx context.Context, clientc chan []Repository, repos []Repository) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	select {
	case <-timeoutCtx.Done():
	case clientc <- repos:
	}
}

func (s AutoRefreshingService) GetAllPRs(ctx context.Context) (repos []Repository, err error) {
	if repos, ok := s.reposFromCache(); ok {
		return repos, nil
	}

	defer func() {
		s.reposToCache(repos)
	}()

	return s.service.GetAllPRs(ctx)
}

func (s AutoRefreshingService) StreamAllPRs(ctx context.Context) (chan []Repository, error) {
	return nil, nil
}

func (s AutoRefreshingService) reposFromCache() ([]Repository, bool) {
	repoBytes, err := s.cache.Get(cacheKey)
	if err != nil {
		return nil, false
	}

	var repos []Repository
	if err := json.Unmarshal(repoBytes, &repos); err != nil {
		return nil, false
	}

	return repos, true
}

func (s AutoRefreshingService) reposToCache(repos []Repository) {
	if repos == nil {
		return
	}

	repoBytes, err := json.Marshal(repos)
	if err != nil {
		return
	}

	_ = s.cache.Set(cacheKey, repoBytes)
}
