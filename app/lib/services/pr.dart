import 'package:app/pb/dashboard/v1/dashboard.pbgrpc.dart';
import 'package:grpc/grpc_or_grpcweb.dart';

class PrService {
  late DashboardServiceClient client;

  PrService({required endpoint}) {
    GrpcOrGrpcWebClientChannel channel =
        GrpcOrGrpcWebClientChannel.toSeparatePorts(
            host: endpoint,
            grpcPort: 8081,
            grpcTransportSecure: false,
            grpcWebPort: 8080,
            grpcWebTransportSecure: false);
    client = DashboardServiceClient(channel);
  }

  Future<List<Repository>> getAllPRs() async {
    var repos = await client.listPullRequests(ListPullRequestsRequest());
    return repos.items
        .map((repo) => Repository(
            name: repo.name,
            pullrequests: repo.pullrequests
                .map((pr) => PR(title: pr.title, url: pr.url))
                .toList()))
        .toList();
  }
}

class PR {
  String title;
  String url;

  PR({required this.title, required this.url});
}

class Repository {
  String name;
  List<PR> pullrequests;

  Repository({required this.name, required this.pullrequests});
}
