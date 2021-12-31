import 'package:app/services/pr.dart';
import 'package:app/ui/pr_card.dart';
import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:url_launcher/url_launcher.dart';

class RepositoryCard extends StatelessWidget {
  final Repository repo;
  const RepositoryCard({Key? key, required this.repo}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Card(
      color: const Color(0xffedf2f7),
      elevation: 2.0,
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
                      icon: const FaIcon(FontAwesomeIcons.github))
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
    );
  }
}
