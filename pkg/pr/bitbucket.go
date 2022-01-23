package pr

import (
	"context"
	"fmt"
	"time"

	bitbucketv1 "github.com/gfleury/go-bitbucket-v1"
)

var _ Provider = (*Bitbucket)(nil)

type Bitbucket struct {
	client *bitbucketv1.APIClient
}

type BitbucketConfig struct {
	// baseURL for the API, path like "/api/1.0/projects/{projectKey}" will be appended.
	BaseURL string
	PAT     string
}

func NewBitbucketProvider(cfg *BitbucketConfig) (*Bitbucket, error) {
	authCtx := context.WithValue(context.Background(), bitbucketv1.ContextAccessToken, cfg.PAT)
	client := bitbucketv1.NewAPIClient(authCtx, bitbucketv1.NewConfiguration(cfg.BaseURL))

	return &Bitbucket{
		client: client,
	}, nil
}

func (b Bitbucket) ListReposForProject(ctx context.Context, project string) ([]string, error) {
	var repoIDs []string

	b.forEachPage(func(opts map[string]interface{}) (*bitbucketv1.APIResponse, error) {
		resp, err := b.client.DefaultApi.GetRepositoriesWithOptions(project, opts)
		if err != nil {
			return nil, err
		}

		repos, err := bitbucketv1.GetRepositoriesResponse(resp)
		if err != nil {
			return nil, err
		}

		for _, repo := range repos {
			repoIDs = append(repoIDs, fmt.Sprintf("%s/%s", project, repo.Slug))
		}

		return resp, err
	})

	return repoIDs, nil
}

// GetPRsForRepo returns all PRs for a given repo.
// repoID is the unique identifier of the repository.
// For GitHub it would be sth like owner/repository.
// For AzureDevops it would be sth like project/repository.
func (b Bitbucket) GetPRsForRepo(ctx context.Context, repoID string) ([]PR, error) {
	project, repo, err := projectRepoFromID(repoID)
	if err != nil {
		return nil, err
	}

	pullrequests, err := b.fetchAllPRs(project, repo)
	if err != nil {
		return nil, err
	}

	prs := make([]PR, 0, len(pullrequests))
	for _, pullrequest := range pullrequests {
		prs = append(prs, PR{
			Title:        pullrequest.Title,
			SourceBranch: pullrequest.FromRef.DisplayID,
			TargetBranch: pullrequest.ToRef.DisplayID,
			User:         pullrequest.Author.User.DisplayName,
			CreatedAt:    time.UnixMilli(pullrequest.CreatedDate),
			Status:       b.statusForPR(pullrequest.State),
			URL:          pullrequest.Links.Self[0].Href,
		})
	}

	return prs, err
}

func (b Bitbucket) fetchAllPRs(project, repo string) ([]bitbucketv1.PullRequest, error) {
	var pullrequests []bitbucketv1.PullRequest

	b.forEachPage(func(opts map[string]interface{}) (*bitbucketv1.APIResponse, error) {
		resp, err := b.client.DefaultApi.GetPullRequestsPage(project, repo, opts)
		if err != nil {
			return nil, err
		}

		pullrequestsPart, err := bitbucketv1.GetPullRequestsResponse(resp)
		if err != nil {
			return nil, err
		}
		pullrequests = append(pullrequests, pullrequestsPart...)

		return resp, err
	})

	return pullrequests, nil
}

// GetURLForRepo returns the web URL for a given repo.
func (b Bitbucket) GetURLForRepo(ctx context.Context, repoID string) (string, error) {
	project, repoName, err := projectRepoFromID(repoID)
	if err != nil {
		return "", err
	}

	resp, err := b.client.DefaultApi.GetRepository(project, repoName)
	if err != nil {
		return "", err
	}

	repo, err := bitbucketv1.GetRepositoryResponse(resp)
	if err != nil {
		return "", err
	}

	return repo.Links.Self[0].Href, nil
}

func (b Bitbucket) statusForPR(status string) PRStatus {
	switch status {
	case "OPEN":
		return PRStatusActive
	case "DECLINED", "MERGED":
		return PRStatusClosed
	default:
		return PRStatusUnspecified
	}
}

func (b Bitbucket) forEachPage(pageFn func(opts map[string]interface{}) (*bitbucketv1.APIResponse, error)) error {
	start := 0
	for {
		resp, err := pageFn(map[string]interface{}{"start": start})
		if err != nil {
			return err
		}

		hasNextPage, nextPage := bitbucketv1.HasNextPage(resp)
		if !hasNextPage {
			break
		}

		start = nextPage
	}

	return nil
}
