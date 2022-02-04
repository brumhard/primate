package pr

import (
	"context"
	"net/http"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

var _ Provider = (*GitHub)(nil)

type GitHub struct {
	client *github.Client
}

type GitHubConfig struct {
	PAT string
}

func NewGitHubProvider(cfg *GitHubConfig) (*GitHub, error) {
	var httpClient *http.Client
	if cfg != nil && cfg.PAT != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: cfg.PAT},
		)
		httpClient = oauth2.NewClient(context.Background(), ts)
	}

	client := github.NewClient(httpClient)

	return &GitHub{
		client: client,
	}, nil
}

func (g GitHub) ListRepositoriesForProject(ctx context.Context, project string) ([]string, error) {
	var repoNames []string

	listFunc := func(ctx context.Context, project string, opts github.ListOptions) ([]*github.Repository, *github.Response, error) {
		return g.client.Repositories.List(ctx, project, &github.RepositoryListOptions{ListOptions: opts})
	}
	// check if it's an org, since then ListByOrg is needed to also list private repos
	if _, _, err := g.client.Organizations.Get(ctx, project); err == nil {
		listFunc = func(ctx context.Context, project string, opts github.ListOptions) ([]*github.Repository, *github.Response, error) {
			return g.client.Repositories.ListByOrg(ctx, project, &github.RepositoryListByOrgOptions{ListOptions: opts})
		}
	}

	err := g.forEachPage(func(opts github.ListOptions) (*github.Response, error) {
		repos, resp, err := listFunc(ctx, project, opts)
		if err != nil {
			return nil, err
		}
		for _, repo := range repos {
			repoNames = append(repoNames, repo.GetName())
		}

		return resp, nil
	})
	if err != nil {
		return nil, err
	}

	return repoNames, nil
}

func (g GitHub) ListPullRequestsForRepository(ctx context.Context, project, repo string) ([]PR, error) {
	var prs []PR
	err := g.forEachPage(func(opts github.ListOptions) (*github.Response, error) {
		pullrequests, resp, err := g.client.PullRequests.List(ctx, project, repo, &github.PullRequestListOptions{ListOptions: opts})
		if err != nil {
			return nil, err
		}

		for _, pullrequest := range pullrequests {
			prs = append(prs, PR{
				Title:        pullrequest.GetTitle(),
				URL:          pullrequest.GetHTMLURL(),
				User:         pullrequest.GetUser().GetLogin(),
				SourceBranch: pullrequest.GetHead().GetRef(),
				TargetBranch: pullrequest.GetBase().GetRef(),
				CreatedAt:    pullrequest.GetCreatedAt(),
				Status:       g.statusForPR(pullrequest),
			})
		}

		return resp, nil
	})
	if err != nil {
		return nil, err
	}

	return prs, nil
}

func (g GitHub) GetRepositoryURL(ctx context.Context, project, repoName string) (string, error) {
	repo, _, err := g.client.Repositories.Get(ctx, project, repoName)
	if err != nil {
		return "", err
	}

	return repo.GetHTMLURL(), nil
}

func (g GitHub) statusForPR(pr *github.PullRequest) Status {
	if pr.GetDraft() {
		return StatusDraft
	}

	switch pr.GetState() {
	case "open":
		return StatusActive
	case "closed":
		return StatusClosed
	default:
		return StatusUnspecified
	}
}

func (g GitHub) forEachPage(pageFn func(opts github.ListOptions) (*github.Response, error)) error {
	opts := github.ListOptions{PerPage: 20}

	for {
		resp, err := pageFn(opts)
		if err != nil {
			return err
		}

		if resp == nil || resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return nil
}
