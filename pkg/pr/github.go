package pr

import (
	"context"
	"errors"
	"net/http"
	"strings"

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

func ownerRepoFromID(repoID string) (string, string, error) {
	split := strings.Split(repoID, "/")
	if len(split) != 2 {
		return "", "", errors.New("malformed repoID")
	}

	return split[0], split[1], nil
}

func (g GitHub) ListReposForProject(ctx context.Context, project string) ([]string, error) {
	var repoIDs []string

	err := forEachPage(func(opts github.ListOptions) (*github.Response, error) {
		repos, resp, err := g.client.Repositories.List(ctx, project, &github.RepositoryListOptions{ListOptions: opts})
		if err != nil {
			return nil, err
		}
		for _, repo := range repos {
			repoIDs = append(repoIDs, repo.GetFullName())
		}

		return resp, nil
	})
	if err != nil {
		return nil, err
	}

	return repoIDs, nil
}

func (g GitHub) GetPRsForRepo(ctx context.Context, repoID string) ([]PR, error) {
	owner, repo, err := ownerRepoFromID(repoID)
	if err != nil {
		return nil, err
	}

	var prs []PR
	err = forEachPage(func(opts github.ListOptions) (*github.Response, error) {
		pullrequests, resp, err := g.client.PullRequests.List(ctx, owner, repo, &github.PullRequestListOptions{ListOptions: opts})
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

func (g GitHub) GetURLForRepo(ctx context.Context, repoID string) (string, error) {
	owner, repoName, err := ownerRepoFromID(repoID)
	if err != nil {
		return "", err
	}

	repo, _, err := g.client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return "", err
	}

	return repo.GetHTMLURL(), nil
}

func (g GitHub) statusForPR(pr *github.PullRequest) PRStatus {
	if pr.GetDraft() {
		return PRStatusDraft
	}

	switch pr.GetState() {
	case "open":
		return PRStatusActive
	case "closed":
		return PRStatusClosed
	default:
		return PRStatusUnspecified
	}
}

func forEachPage(pageFn func(opts github.ListOptions) (*github.Response, error)) error {
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
