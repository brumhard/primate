package pr

import "context"

var _ Provider = (*Bitbucket)(nil)

type Bitbucket struct {
}

// GetPRsForRepo returns all PRs for a given repo.
// repoID is the unique identifier of the repository.
// For GitHub it would be sth like owner/repository.
// For AzureDevops it would be sth like project/repository.
func (b Bitbucket) GetPRsForRepo(ctx context.Context, repoID string) ([]PR, error) {
	panic("not implemented") // TODO: Implement
}

// GetURLForRepo returns the web URL for a given repo.
func (b Bitbucket) GetURLForRepo(ctx context.Context, repoID string) (string, error) {
	panic("not implemented") // TODO: Implement
}
