import 'dart:html';

import 'package:app/services/pr.dart';
import 'package:app/ui/pr_card.dart';
import 'package:flutter/material.dart';

class RepositoryCard extends StatelessWidget {
  String name;
  List<PR> prs;
  RepositoryCard({Key? key, required this.name, required this.prs})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Column(
        children: [
          SelectableText(name, style: Theme.of(context).textTheme.headline2),
          ...prs.map((pr) => PRCard(pr: pr)).toList()
        ],
      ),
    );
  }
}
