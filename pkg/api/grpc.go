package api

import (
	"context"
	"time"

	dashboardv1 "github.com/brumhard/pr-dashboard/pkg/pb/dashboard/v1"
	"github.com/brumhard/pr-dashboard/pkg/pr"
)

var _ dashboardv1.DashboardServiceServer = (*GRPC)(nil)

type GRPC struct {
	dashboardv1.UnimplementedDashboardServiceServer
	service pr.Service
}

func NewGRPC(service pr.Service) *GRPC {
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
				Title:        pullrequest.Title,
				Url:          pullrequest.URL,
				User:         pullrequest.User,
				SourceBranch: pullrequest.SourceBranch,
				TargetBranch: pullrequest.TargetBranch,
				CreatedAt:    pullrequest.CreatedAt.Format(time.RFC3339),
				Status:       dashboardv1.PullRequest_Status(pullrequest.Status),
			})
		}

		grpcRepos = append(grpcRepos, &dashboardv1.Repository{
			Name:         repo.Name,
			Url:          repo.URL,
			Pullrequests: grpcPRs,
		})
	}
	return &dashboardv1.ListPullRequestsResponse{
		Items: grpcRepos,
	}, nil
}
