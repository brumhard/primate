package pr

import (
	"context"
	"regexp"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

var ErrUnknownProvider = errors.New("unknown provider")

type Provider interface {
	// ListRepositoriesForProject fetches all repositories in a project
	// (could be called sth else in each provider, e.g. it would be owner/org in github).
	// It returns a list of repository names.
	ListRepositoriesForProject(ctx context.Context, project string) ([]string, error)
	// ListPullRequestsForRepository returns all PRs for a given project and repo.
	// For GitHub the project would be an organization or user.
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

type StreamerService interface {
	Service
	StreamAllPRs(ctx context.Context) (chan []Repository, error)
}

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
	// TODO: run concurrently to save time on slow responses
	for _, repoID := range s.repos {
		project, repoName, err := projectRepoFromID(repoID)
		if err != nil {
			return nil, err
		}

		repoRegex, err := regexp.Compile(repoName)
		if err != nil {
			// TODO: if no regex can be created, still try to fetch the repo with given string
			return nil, err
		}

		reposInProject, err := s.provider.ListRepositoriesForProject(ctx, project)
		if err != nil {
			return nil, err
		}

		for _, repo := range reposInProject {
			if !repoRegex.MatchString(repo) {
				continue
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

	}

	return repos, nil
}
