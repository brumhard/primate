package pr

import (
	"context"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

var ErrUnknownProvider = errors.New("unknown provider")

type Provider interface {
	// ListRepositoriesForProject fetches all repositories in a project
	// (could be called sth else in each provider, e.g. it would be owner/org in github).
	// It returns a list of repoIDs that should be valid input for GetPRsForRepo.
	ListRepositoriesForProject(ctx context.Context, project string) ([]string, error)
	// ListPullRequestsForRepository returns all PRs for a given repo.
	// repoID is the unique identifier of the repository.
	// For GitHub it would be sth like owner/repository.
	// For AzureDevops it would be sth like project/repository.
	ListPullRequestsForRepository(ctx context.Context, project, repo string) ([]PR, error)
	// GetRepositoryURL returns the web URL for a given repo.
	GetRepositoryURL(ctx context.Context, project, repo string) (string, error)
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

func projectRepoFromID(repoID string) (project, repo string, err error) {
	split := strings.Split(repoID, "/")
	if len(split) != 2 {
		return "", "", errors.New("malformed repoID")
	}

	return split[0], split[1], nil
}

func (s SingleProviderService) GetAllPRs(ctx context.Context) ([]Repository, error) {
	repos := make([]Repository, 0, len(s.repos))
	for _, repoID := range s.repos {
		project, repo, err := projectRepoFromID(repoID)
		if err != nil {
			return nil, err
		}

		prs, err := s.provider.ListPullRequestsForRepository(ctx, project, repo)
		if err != nil {
			return nil, err
		}

		repoURL, err := s.provider.GetRepositoryURL(ctx, project, repo)
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
