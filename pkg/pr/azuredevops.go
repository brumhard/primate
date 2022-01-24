package pr

import (
	"context"
	"fmt"
	"strings"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/git"
)

var _ Provider = (*AzureDevops)(nil)

type AzureDevops struct {
	conn *azuredevops.Connection
}

type AzureDevopsConfig struct {
	OrganizationURL string
	PAT             string
}

func NewAzureDevopsProvider(cfg *AzureDevopsConfig) (*AzureDevops, error) {
	// TODO: implement cfg validation

	// Create a connection to your organization
	connection := azuredevops.NewPatConnection(cfg.OrganizationURL, cfg.PAT)

	return &AzureDevops{conn: connection}, nil
}

func (ad AzureDevops) ListRepositoriesForProject(ctx context.Context, project string) ([]string, error) {
	gitClient, err := git.NewClient(ctx, ad.conn)
	if err != nil {
		return nil, err
	}

	repos, err := gitClient.GetRepositories(ctx, git.GetRepositoriesArgs{Project: &project})
	if err != nil {
		return nil, err
	}

	repoIDs := make([]string, 0, len(*repos))
	for _, repo := range *repos {
		repoIDs = append(repoIDs, fmt.Sprintf("%s/%s", project, *repo.Name))
	}

	return repoIDs, nil
}

func (ad AzureDevops) ListPullRequestsForRepository(ctx context.Context, project, repo string) ([]PR, error) {
	gitClient, err := git.NewClient(ctx, ad.conn)
	if err != nil {
		return nil, err
	}

	pullrequests, err := gitClient.GetPullRequests(ctx, git.GetPullRequestsArgs{
		RepositoryId:   &repo,
		Project:        &project,
		SearchCriteria: &git.GitPullRequestSearchCriteria{},
	})
	if err != nil {
		return nil, err
	}

	refPrefix := "refs/heads/"

	prs := make([]PR, 0, len(*pullrequests))
	for _, pullrequest := range *pullrequests {
		prs = append(prs, PR{
			Title:        *pullrequest.Title,
			URL:          webURLFromPR(pullrequest),
			User:         *pullrequest.CreatedBy.DisplayName,
			SourceBranch: strings.TrimPrefix(*pullrequest.SourceRefName, refPrefix),
			TargetBranch: strings.TrimPrefix(*pullrequest.TargetRefName, refPrefix),
			CreatedAt:    pullrequest.CreationDate.Time,
			Status:       ad.statusForPR(pullrequest),
		})
	}

	return prs, nil
}

func (ad AzureDevops) GetRepositoryURL(ctx context.Context, project, repoName string) (string, error) {
	gitClient, err := git.NewClient(ctx, ad.conn)
	if err != nil {
		return "", err
	}

	repo, err := gitClient.GetRepository(ctx, git.GetRepositoryArgs{
		RepositoryId: &repoName,
		Project:      &project,
	})
	if err != nil {
		return "", err
	}

	return *repo.WebUrl, nil
}

func (ad AzureDevops) statusForPR(pr git.GitPullRequest) Status {
	if pr.IsDraft != nil && *pr.IsDraft {
		return StatusDraft
	}

	if pr.Status == nil {
		return StatusUnspecified
	}

	switch *pr.Status {
	case "active":
		return StatusActive
	case "completed":
		return StatusClosed
	default:
		return StatusUnspecified
	}
}

func webURLFromPR(pullrequest git.GitPullRequest) string {
	if pullrequest.Url == nil {
		return ""
	}

	replacer := strings.NewReplacer(
		"_apis/git/repositories", "_git",
		"/pullRequests", "/pullrequest",
	)
	return replacer.Replace(*pullrequest.Url)
}
