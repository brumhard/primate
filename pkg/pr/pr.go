package pr

import "time"

type Repository struct {
	Name         string
	URL          string
	PullRequests []PR
}

type Status int

const (
	StatusUnspecified Status = iota
	StatusDraft
	StatusActive
	StatusClosed
)

type PR struct {
	Title        string
	URL          string
	User         string
	SourceBranch string
	TargetBranch string
	CreatedAt    time.Time
	Status       Status
}
