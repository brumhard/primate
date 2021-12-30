import 'package:app/services/pr.dart';
import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:url_launcher/url_launcher.dart';

class PRCard extends StatelessWidget {
  final PR pr;
  const PRCard({Key? key, required this.pr}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Card(
      elevation: 2.0,
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            SelectableText(pr.title),
            IconButton(
              icon: const FaIcon(
                FontAwesomeIcons.gitAlt,
              ),
              onPressed: () async {
                if (!await canLaunch(pr.url)) {
                  throw 'Could not launch ${pr.url}';
                }
                await launch(pr.url);
              },
            )
          ],
        ),
      ),
    );
  }
}
