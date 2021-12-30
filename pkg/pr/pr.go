package pr

type PR struct {
	Title string
	URL   string
	User  string
}

type Repository struct {
	Name         string
	PullRequests []PR
}
