package pr

import "time"

type Repository struct {
	Name         string
	URL          string
	PullRequests []PR
}

type PRStatus int

const (
	PRStatusUnspecified PRStatus = iota
	PrStatusDraft
	PrStatusActive
	PrStatusClosed
)

type PR struct {
	Title        string
	URL          string
	User         string
	SourceBranch string
	TargetBranch string
	CreatedAt    time.Time
	Status       PRStatus
}
