import 'dart:async';

import 'package:primate/pb/dashboard/v1/dashboard.pbgrpc.dart';
import 'package:flutter/material.dart';
import 'package:grpc/grpc_or_grpcweb.dart';

class PrService {
  late DashboardServiceClient client;
  final StreamController<List<Repository>?> _controller =
      StreamController.broadcast();
  final ValueNotifier<bool> _isLoading = ValueNotifier(false);

  PrService({required endpoint}) {
    GrpcOrGrpcWebClientChannel channel =
        GrpcOrGrpcWebClientChannel.toSeparatePorts(
            host: endpoint,
            grpcPort: 8081,
            grpcTransportSecure: false,
            grpcWebPort: 8080,
            grpcWebTransportSecure: false);
    client = DashboardServiceClient(channel);

    _controller.onListen = () => loadPRs();
  }

  Stream<List<Repository>?> get stream => _controller.stream;
  ValueNotifier<bool> get isLoading => _isLoading;

  void dispose() {
    _controller.close();
  }

  void loadPRs() async {
    // set loading state
    _isLoading.value = true;
    ListPullRequestsResponse reposResponse;
    try {
      reposResponse = await client.listPullRequests(ListPullRequestsRequest());
    } catch (e) {
      _controller.addError(e);
      _isLoading.value = false;
      return;
    }

    var repos = reposResponse.items
        .map((repo) => Repository(
            name: repo.name,
            url: repo.url,
            pullrequests: repo.pullrequests
                .map((pr) => PR(
                      title: pr.title,
                      url: pr.url,
                      user: pr.user,
                      sourceBranch: pr.sourceBranch,
                      targetBranch: pr.targetBranch,
                      created: DateTime.parse(pr.createdAt),
                      status: _prStatusToName(pr.status),
                    ))
                .toList()))
        .toList();
    _controller.add(repos);
    _isLoading.value = false;
  }
}

String _prStatusToName(PullRequest_Status status) {
  switch (status) {
    case PullRequest_Status.STATUS_ACTIVE:
      return "active";
    case PullRequest_Status.STATUS_DRAFT:
      return "draft";
    case PullRequest_Status.STATUS_CLOSED:
      return "closed";
    default:
      return "unspecified";
  }
}

class PR {
  String title;
  String url;
  String user;
  String sourceBranch;
  String targetBranch;
  DateTime created;
  String status;

  PR(
      {required this.title,
      required this.url,
      required this.user,
      required this.sourceBranch,
      required this.targetBranch,
      required this.created,
      required this.status});
}

class Repository {
  String name;
  String url;
  List<PR> pullrequests;

  Repository(
      {required this.name, required this.url, required this.pullrequests});
}
