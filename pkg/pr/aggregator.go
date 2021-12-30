package pr

import "context"

type Aggregator struct {
	services []Service
}

func NewAggregator(configs []ProviderConfiguration) (*Aggregator, error) {
	services := make([]Service, 0, len(configs))
	for _, config := range configs {
		svc, err := NewService(config)
		if err != nil {
			return nil, err
		}

		services = append(services, *svc)
	}

	return &Aggregator{services: services}, nil
}

func (a Aggregator) GetAllPRs(ctx context.Context) ([]Repository, error) {
	var allRepos []Repository
	for _, svc := range a.services {
		repos, err := svc.GetAllPRs(ctx)
		if err != nil {
			return nil, err
		}

		allRepos = append(allRepos, repos...)
	}

	return allRepos, nil
}
