package pr

import (
	"context"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

var ErrUnknownProvider = errors.New("unknown provider")

type Provider interface {
	// ListReposForProject fetches all repositories in a project
	// (could be called sth else in each provider, e.g. it would be owner/org in github).
	// It returns a list of repoIDs that should be valid input for GetPRsForRepo.
	ListReposForProject(ctx context.Context, project string) ([]string, error)
	// GetPRsForRepo returns all PRs for a given repo.
	// repoID is the unique identifier of the repository.
	// For GitHub it would be sth like owner/repository.
	// For AzureDevops it would be sth like project/repository.
	GetPRsForRepo(ctx context.Context, repoID string) ([]PR, error)
	// GetURLForRepo returns the web URL for a given repo.
	GetURLForRepo(ctx context.Context, repoID string) (string, error)
}

type ProviderType string

const (
	ProviderTypeGitHub      = "github"
	ProviderTypeAzureDevops = "azuredevops"
	ProviderTypeBitbucketV1 = "bitbucketv1"
)

type Service interface {
	GetAllPRs(ctx context.Context) ([]Repository, error)
}

var _ Service = (*SingleProviderService)(nil)

type SingleProviderService struct {
	repos    []string
	provider Provider
}

type ProviderConfiguration struct {
	Repositories []string
	ProviderType ProviderType
	ExtraConfig  map[string]interface{}
}

func NewSingleProviderService(config ProviderConfiguration) (*SingleProviderService, error) {
	provider, err := providerForType(config.ProviderType, config.ExtraConfig)
	if err != nil {
		return nil, err
	}

	return &SingleProviderService{repos: config.Repositories, provider: provider}, nil
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
	case ProviderTypeBitbucketV1:
		var config BitbucketConfig
		if err := mapstructure.Decode(cfg, &config); err != nil {
			return nil, err
		}

		return NewBitbucketProvider(&config)
	}
	return nil, errors.Wrap(ErrUnknownProvider, string(providerType))
}

// TODO: include functionality to include repo url in response
func (s SingleProviderService) GetAllPRs(ctx context.Context) ([]Repository, error) {
	repos := make([]Repository, 0, len(s.repos))
	for _, repo := range s.repos {
		prs, err := s.provider.GetPRsForRepo(ctx, repo)
		if err != nil {
			return nil, err
		}

		repoURL, err := s.provider.GetURLForRepo(ctx, repo)
		if err != nil {
			return nil, err
		}

		repos = append(repos, Repository{
			Name:         repo,
			URL:          repoURL,
			PullRequests: prs,
		})
	}

	return repos, nil
}
