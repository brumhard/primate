package pr

import (
	"context"
	"errors"
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

func projectRepoFromID(repoID string) (string, string, error) {
	split := strings.Split(repoID, "/")
	if len(split) != 2 {
		return "", "", errors.New("malformed repoID")
	}

	return split[0], split[1], nil
}

func (ad AzureDevops) GetPRsForRepo(ctx context.Context, repoID string) ([]PR, error) {
	gitClient, err := git.NewClient(ctx, ad.conn)
	if err != nil {
		return nil, err
	}

	project, repo, err := projectRepoFromID(repoID)
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

	prs := make([]PR, 0, len(*pullrequests))
	for _, pullrequest := range *pullrequests {
		prs = append(prs, PR{
			Title: *pullrequest.Title,
			URL:   webURLFromPR(pullrequest),
		})
	}

	return prs, nil
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