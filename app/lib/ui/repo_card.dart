import 'package:app/services/pr.dart';
import 'package:app/ui/pr_card.dart';
import 'package:flutter/material.dart';

class RepositoryCard extends StatelessWidget {
  final Repository repo;
  const RepositoryCard({Key? key, required this.repo}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Card(
      elevation: 2.0,
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: SelectableText(repo.name,
                  style: Theme.of(context).textTheme.headline5),
            ),
            ...repo.pullrequests.map((pr) => PRCard(pr: pr)).toList()
          ],
        ),
      ),
    );
  }
}
