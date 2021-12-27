package pr

import (
	"context"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

var ErrUnknownProvider = errors.New("unknown provider")

type Provider interface {
	// GetPRsForRepo returns all PRs for a given repo.
	// repoID is the unique identifier of the repository.
	// For GitHub it would be sth like owner/repository.
	// For AzureDevops it would be sth like project/repository.
	GetPRsForRepo(ctx context.Context, repoID string) ([]PR, error)
}

type ProviderType string

const (
	ProviderTypeGitHub      = "github"
	ProviderTypeAzureDevops = "azuredevops"
)

type ProviderConfiguration struct {
	Repositories []string
	ProviderType ProviderType
	ExtraConfig  map[string]interface{}
}

func NewService(config ProviderConfiguration) (*Service, error) {
	provider, err := providerForType(config.ProviderType, config.ExtraConfig)
	if err != nil {
		return nil, err
	}

	return &Service{repos: config.Repositories, provider: provider}, nil
}

func providerForType(providerType ProviderType, cfg map[string]interface{}) (Provider, error) {
	switch providerType {
	case ProviderTypeAzureDevops:
		var config AzureDevopsConfig
		if err := mapstructure.Decode(cfg, &config); err != nil {
			return nil, err
		}

		return NewAzureDevopsProvider(&config)
	case ProviderTypeGitHub:
		var config GitHubConfig
		if err := mapstructure.Decode(cfg, &config); err != nil {
			return nil, err
		}

		return NewGitHubProvider(&config)
	}
	return nil, errors.Wrap(ErrUnknownProvider, string(providerType))
}

type Service struct {
	repos    []string
	provider Provider
}

func (s Service) GetAllPRs(ctx context.Context) ([]PR, error) {
	allPRs := make([]PR, 0, len(s.repos))
	for _, repo := range s.repos {
		prs, err := s.provider.GetPRsForRepo(ctx, repo)
		if err != nil {
			return nil, err
		}

		allPRs = append(prs, prs...)
	}

	return allPRs, nil
}
