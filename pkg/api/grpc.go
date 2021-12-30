package api

import (
	"context"

	dashboardv1 "github.com/brumhard/pr-dashboard/pkg/pb/dashboard/v1"
	"github.com/brumhard/pr-dashboard/pkg/pr"
)

var _ dashboardv1.DashboardServiceServer = (*GRPC)(nil)

type GRPC struct {
	dashboardv1.UnimplementedDashboardServiceServer
	service *pr.Aggregator
}

func NewGRPC(service *pr.Aggregator) *GRPC {
	return &GRPC{service: service}
}

func (g *GRPC) ListPullRequests(ctx context.Context, request *dashboardv1.ListPullRequestsRequest) (*dashboardv1.ListPullRequestsResponse, error) {
	repos, err := g.service.GetAllPRs(ctx)
	if err != nil {
		return nil, err
	}

	grpcRepos := make([]*dashboardv1.Repository, 0, len(repos))
	for _, repo := range repos {
		grpcPRs := make([]*dashboardv1.PullRequest, 0, len(repo.PullRequests))
		for _, pullrequest := range repo.PullRequests {
			grpcPRs = append(grpcPRs, &dashboardv1.PullRequest{
				Title: pullrequest.Title,
				Url:   pullrequest.URL,
				User:  pullrequest.User,
			})
		}

		grpcRepos = append(grpcRepos, &dashboardv1.Repository{
			Name:         repo.Name,
			Pullrequests: grpcPRs,
		})
	}
	return &dashboardv1.ListPullRequestsResponse{
		Items: grpcRepos,
	}, nil
}
