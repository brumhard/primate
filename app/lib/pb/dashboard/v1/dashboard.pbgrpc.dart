///
//  Generated code. Do not modify.
//  source: dashboard/v1/dashboard.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'dashboard.pb.dart' as $0;
export 'dashboard.pb.dart';

class DashboardServiceClient extends $grpc.Client {
  static final _$listPullRequests = $grpc.ClientMethod<
          $0.ListPullRequestsRequest, $0.ListPullRequestsResponse>(
      '/dashboard.v1.DashboardService/ListPullRequests',
      ($0.ListPullRequestsRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.ListPullRequestsResponse.fromBuffer(value));
  static final _$streamPullRequests = $grpc.ClientMethod<
          $0.StreamPullRequestsRequest, $0.ListPullRequestsResponse>(
      '/dashboard.v1.DashboardService/StreamPullRequests',
      ($0.StreamPullRequestsRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.ListPullRequestsResponse.fromBuffer(value));

  DashboardServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.ListPullRequestsResponse> listPullRequests(
      $0.ListPullRequestsRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$listPullRequests, request, options: options);
  }

  $grpc.ResponseStream<$0.ListPullRequestsResponse> streamPullRequests(
      $0.StreamPullRequestsRequest request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(
        _$streamPullRequests, $async.Stream.fromIterable([request]),
        options: options);
  }
}

abstract class DashboardServiceBase extends $grpc.Service {
  $core.String get $name => 'dashboard.v1.DashboardService';

  DashboardServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.ListPullRequestsRequest,
            $0.ListPullRequestsResponse>(
        'ListPullRequests',
        listPullRequests_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.ListPullRequestsRequest.fromBuffer(value),
        ($0.ListPullRequestsResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.StreamPullRequestsRequest,
            $0.ListPullRequestsResponse>(
        'StreamPullRequests',
        streamPullRequests_Pre,
        false,
        true,
        ($core.List<$core.int> value) =>
            $0.StreamPullRequestsRequest.fromBuffer(value),
        ($0.ListPullRequestsResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.ListPullRequestsResponse> listPullRequests_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.ListPullRequestsRequest> request) async {
    return listPullRequests(call, await request);
  }

  $async.Stream<$0.ListPullRequestsResponse> streamPullRequests_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.StreamPullRequestsRequest> request) async* {
    yield* streamPullRequests(call, await request);
  }

  $async.Future<$0.ListPullRequestsResponse> listPullRequests(
      $grpc.ServiceCall call, $0.ListPullRequestsRequest request);
  $async.Stream<$0.ListPullRequestsResponse> streamPullRequests(
      $grpc.ServiceCall call, $0.StreamPullRequestsRequest request);
}
