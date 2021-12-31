import 'dart:convert';

import 'package:app/services/pr.dart';
import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:intl/intl.dart';
import 'package:url_launcher/link.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:crypto/crypto.dart';

class PRCard extends StatelessWidget {
  final PR pr;
  const PRCard({Key? key, required this.pr}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    // TODO: use Link widget for correct context menu as soon as https://github.com/flutter/flutter/issues/91881 is closed

    // hash user to not expose personal info
    var userHash = sha1.convert(utf8.encode(pr.user));

    return InkWell(
      onTap: () async {
        if (!await canLaunch(pr.url)) {
          throw 'Could not launch ${pr.url}';
        }
        await launch(pr.url);
      },
      borderRadius: BorderRadius.circular(5),
      child: Card(
        elevation: 2.0,
        child: Padding(
          padding: const EdgeInsets.all(8.0),
          child: Row(mainAxisAlignment: MainAxisAlignment.start, children: [
            Tooltip(
              message: pr.user,
              child: CircleAvatar(
                backgroundColor: Colors.transparent,
                backgroundImage: NetworkImage(
                    "https://avatars.dicebear.com/api/pixel-art/$userHash.png"),
              ),
            ),
            Expanded(
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  Padding(
                    padding: const EdgeInsets.only(left: 8.0),
                    child: Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text(
                          pr.title,
                          overflow: TextOverflow.ellipsis,
                          style: const TextStyle(fontWeight: FontWeight.bold),
                        ),
                        Text(
                          "${pr.sourceBranch} \u{2192} ${pr.targetBranch}",
                          overflow: TextOverflow.ellipsis,
                          style: Theme.of(context).textTheme.overline,
                        ),
                      ],
                    ),
                  ),
                  const Padding(
                    padding: EdgeInsets.all(10.0),
                    child: FaIcon(
                      FontAwesomeIcons.codeBranch,
                    ),
                  ),
                ],
              ),
            ),
          ]),
        ),
      ),
    );
  }
}
