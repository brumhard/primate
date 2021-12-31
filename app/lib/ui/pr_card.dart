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
            Flexible(
              child: Row(
                children: [
                  Tooltip(
                    message: pr.user,
                    child: CircleAvatar(
                      backgroundColor: Colors.transparent,
                      backgroundImage: NetworkImage(
                          "https://avatars.dicebear.com/api/pixel-art/${pr.user}.png"),
                    ),
                  ),
                  Flexible(
                    child: Padding(
                      padding: const EdgeInsets.only(left: 8.0),
                      child: Text(
                        pr.title,
                        overflow: TextOverflow.ellipsis,
                      ),
                    ),
                  ),
                ],
              ),
            ),
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
