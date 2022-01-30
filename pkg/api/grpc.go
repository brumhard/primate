package api

import (
	"context"
	"time"

	dashboardv1 "github.com/brumhard/primate/pkg/pb/dashboard/v1"
	"github.com/brumhard/primate/pkg/pr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ dashboardv1.DashboardServiceServer = (*GRPC)(nil)

// TODO: add proper error codes
type GRPC struct {
	dashboardv1.UnimplementedDashboardServiceServer
	service pr.Service
}

func NewGRPC(service pr.Service) *GRPC {
	return &GRPC{service: service}
}

func (g *GRPC) ListPullRequests(
	ctx context.Context, request *dashboardv1.ListPullRequestsRequest,
) (*dashboardv1.ListPullRequestsResponse, error) {
	repos, err := g.service.GetAllPRs(ctx)
	if err != nil {
		return nil, err
	}

	return &dashboardv1.ListPullRequestsResponse{
		Items: castRepos(repos),
	}, nil
}

func (g *GRPC) StreamPullRequests(request *dashboardv1.StreamPullRequestsRequest, server dashboardv1.DashboardService_StreamPullRequestsServer) error {
	// TODO: check if service implements StreamerService
	// if so call Stream method
	// if not return unimplemented
	streamer, ok := g.service.(pr.StreamerService)
	if !ok {
		return status.Errorf(codes.Unimplemented, "no implemented for the currently running service")
	}

	repoc, err := streamer.StreamAllPRs(server.Context())
	if err != nil {
		// TODO: do not send error in case the context was closed
		return err
	}

	for repos := range repoc {
		server.Send(&dashboardv1.ListPullRequestsResponse{Items: castRepos(repos)})
	}
	return nil
}

func castRepos(repos []pr.Repository) []*dashboardv1.Repository {
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
	return grpcRepos
}
