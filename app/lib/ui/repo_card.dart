import 'package:app/services/pr.dart';
import 'package:app/ui/pr_card.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:url_launcher/url_launcher.dart';

class RepositoryCard extends StatelessWidget {
  final Repository repo;
  const RepositoryCard({Key? key, required this.repo}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    repo.pullrequests.sort((a, b) => b.created.compareTo(a.created));

    return Opacity(
      opacity: repo.pullrequests.isNotEmpty ? 1 : 0.5,
      child: Card(
        color: Theme.of(context).dialogBackgroundColor,
        elevation: 2.0,
        shadowColor: Colors.grey,
        child: Padding(
          padding: const EdgeInsets.all(8.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Padding(
                padding: const EdgeInsets.all(8.0),
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Flexible(
                      child: Text(repo.name,
                          overflow: TextOverflow.ellipsis,
                          style: Theme.of(context).textTheme.headline5),
                    ),
                    IconButton(
                      onPressed: () async {
                        if (!await canLaunch(repo.url)) {
                          throw 'Could not launch ${repo.url}';
                        }
                        await launch(repo.url);
                      },
                      icon: SvgPicture.network(
                        iconForRepoUrl(repo.url),
                        color: Theme.of(context).iconTheme.color,
                        height: 30,
                      ),
                      iconSize: 30,
                    ),
                  ],
                ),
              ),
              ...repo.pullrequests
                  .map((pr) => Row(
                        children: [
                          Flexible(flex: 1, child: Container()),
                          Flexible(flex: 5, child: PRCard(pr: pr)),
                        ],
                      ))
                  .toList()
            ],
          ),
        ),
      ),
    );
  }
}

String iconForRepoUrl(String url) {
  if (url.contains("github.com")) {
    return "https://raw.githubusercontent.com/Templarian/MaterialDesign/master/svg/microsoft-azure-devops.svg";
  }
  if (url.contains("dev.azure.com")) {
    return "https://raw.githubusercontent.com/Templarian/MaterialDesign/master/svg/github.svg";
  }
  return "https://raw.githubusercontent.com/Templarian/MaterialDesign/master/svg/git.svg";
}
