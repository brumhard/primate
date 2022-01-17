package pr

import "context"

var _ Service = (*AggregatorService)(nil)

type AggregatorService struct {
	services []SingleProviderService
}

func NewAggregatorService(configs []ProviderConfiguration) (*AggregatorService, error) {
	services := make([]SingleProviderService, 0, len(configs))
	for _, config := range configs {
		svc, err := NewSingleProviderService(config)
		if err != nil {
			return nil, err
		}

		services = append(services, *svc)
	}

	return &AggregatorService{services: services}, nil
}

func (a AggregatorService) GetAllPRs(ctx context.Context) ([]Repository, error) {
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
