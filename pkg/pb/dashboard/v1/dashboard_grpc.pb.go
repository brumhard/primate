// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: dashboard/v1/dashboard.proto

package dashboard

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DashboardServiceClient is the client API for DashboardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DashboardServiceClient interface {
	ListPullRequests(ctx context.Context, in *ListPullRequestsRequest, opts ...grpc.CallOption) (*ListPullRequestsResponse, error)
	StreamPullRequests(ctx context.Context, in *StreamPullRequestsRequest, opts ...grpc.CallOption) (DashboardService_StreamPullRequestsClient, error)
}

type dashboardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDashboardServiceClient(cc grpc.ClientConnInterface) DashboardServiceClient {
	return &dashboardServiceClient{cc}
}

func (c *dashboardServiceClient) ListPullRequests(ctx context.Context, in *ListPullRequestsRequest, opts ...grpc.CallOption) (*ListPullRequestsResponse, error) {
	out := new(ListPullRequestsResponse)
	err := c.cc.Invoke(ctx, "/dashboard.v1.DashboardService/ListPullRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardServiceClient) StreamPullRequests(ctx context.Context, in *StreamPullRequestsRequest, opts ...grpc.CallOption) (DashboardService_StreamPullRequestsClient, error) {
	stream, err := c.cc.NewStream(ctx, &DashboardService_ServiceDesc.Streams[0], "/dashboard.v1.DashboardService/StreamPullRequests", opts...)
	if err != nil {
		return nil, err
	}
	x := &dashboardServiceStreamPullRequestsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DashboardService_StreamPullRequestsClient interface {
	Recv() (*ListPullRequestsResponse, error)
	grpc.ClientStream
}

type dashboardServiceStreamPullRequestsClient struct {
	grpc.ClientStream
}

func (x *dashboardServiceStreamPullRequestsClient) Recv() (*ListPullRequestsResponse, error) {
	m := new(ListPullRequestsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DashboardServiceServer is the server API for DashboardService service.
// All implementations must embed UnimplementedDashboardServiceServer
// for forward compatibility
type DashboardServiceServer interface {
	ListPullRequests(context.Context, *ListPullRequestsRequest) (*ListPullRequestsResponse, error)
	StreamPullRequests(*StreamPullRequestsRequest, DashboardService_StreamPullRequestsServer) error
	mustEmbedUnimplementedDashboardServiceServer()
}

// UnimplementedDashboardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDashboardServiceServer struct {
}

func (UnimplementedDashboardServiceServer) ListPullRequests(context.Context, *ListPullRequestsRequest) (*ListPullRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPullRequests not implemented")
}
func (UnimplementedDashboardServiceServer) StreamPullRequests(*StreamPullRequestsRequest, DashboardService_StreamPullRequestsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamPullRequests not implemented")
}
func (UnimplementedDashboardServiceServer) mustEmbedUnimplementedDashboardServiceServer() {}

// UnsafeDashboardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DashboardServiceServer will
// result in compilation errors.
type UnsafeDashboardServiceServer interface {
	mustEmbedUnimplementedDashboardServiceServer()
}

func RegisterDashboardServiceServer(s grpc.ServiceRegistrar, srv DashboardServiceServer) {
	s.RegisterService(&DashboardService_ServiceDesc, srv)
}

func _DashboardService_ListPullRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPullRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).ListPullRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dashboard.v1.DashboardService/ListPullRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).ListPullRequests(ctx, req.(*ListPullRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardService_StreamPullRequests_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamPullRequestsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DashboardServiceServer).StreamPullRequests(m, &dashboardServiceStreamPullRequestsServer{stream})
}

type DashboardService_StreamPullRequestsServer interface {
	Send(*ListPullRequestsResponse) error
	grpc.ServerStream
}

type dashboardServiceStreamPullRequestsServer struct {
	grpc.ServerStream
}

func (x *dashboardServiceStreamPullRequestsServer) Send(m *ListPullRequestsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// DashboardService_ServiceDesc is the grpc.ServiceDesc for DashboardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DashboardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dashboard.v1.DashboardService",
	HandlerType: (*DashboardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPullRequests",
			Handler:    _DashboardService_ListPullRequests_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamPullRequests",
			Handler:       _DashboardService_StreamPullRequests_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dashboard/v1/dashboard.proto",
}
