package pr

import (
	"context"
	"encoding/json"
	"time"

	"github.com/allegro/bigcache/v3"
)

const cacheKey = "repos"

var _ Service = (*CachedService)(nil)

type CachedService struct {
	service Service
	cache   *bigcache.BigCache
}

func WithCache(service Service) (Service, error) {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(time.Minute))
	if err != nil {
		return nil, err
	}

	return &CachedService{service: service, cache: cache}, nil
}

func (c CachedService) GetAllPRs(ctx context.Context) (repos []Repository, err error) {
	if repos, ok := c.reposFromCache(); ok {
		return repos, nil
	}

	defer func() {
		c.reposToCache(repos)
	}()

	return c.service.GetAllPRs(ctx)
}

func (c CachedService) reposFromCache() ([]Repository, bool) {
	repoBytes, err := c.cache.Get(cacheKey)
	if err != nil {
		return nil, false
	}

	var repos []Repository
	if err := json.Unmarshal(repoBytes, &repos); err != nil {
		return nil, false
	}

	return repos, true
}

func (c CachedService) reposToCache(repos []Repository) {
	if repos == nil {
		return
	}

	repoBytes, err := json.Marshal(repos)
	if err != nil {
		return
	}

	_ = c.cache.Set(cacheKey, repoBytes)
}
