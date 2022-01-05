import 'dart:convert';

import 'package:app/services/pr.dart';
import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:intl/intl.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:crypto/crypto.dart';

class PRCard extends StatefulWidget {
  final PR pr;
  const PRCard({Key? key, required this.pr}) : super(key: key);

  @override
  State<PRCard> createState() => _PRCardState();
}

class _PRCardState extends State<PRCard> {
  bool _hovered = false;

  @override
  Widget build(BuildContext context) {
    // TODO: use Link widget for correct context menu as soon as https://github.com/flutter/flutter/issues/91881 is closed
    return InkWell(
      radius: 3,
      onTap: () async {
        if (!await canLaunch(widget.pr.url)) {
          throw 'Could not launch ${widget.pr.url}';
        }
        await launch(widget.pr.url);
      },
      onHover: (sth) {
        setState(() {
          _hovered = sth;
        });
      },
      borderRadius: BorderRadius.circular(4),
      child: Padding(
        padding: const EdgeInsets.all(4),
        child: AnimatedContainer(
          duration: const Duration(milliseconds: 10),
          decoration: BoxDecoration(
            boxShadow: const [
              BoxShadow(color: Colors.grey, offset: Offset(3, 3), blurRadius: 3)
            ],
            borderRadius: BorderRadius.circular(4),
            color: _hovered ? null : Colors.white,
            gradient: _hovered
                ? const LinearGradient(
                    begin: Alignment.topLeft,
                    end: Alignment.bottomRight,
                    colors: [Color(0xFFdd5e89), Color(0xFFf7bb97)],
                  )
                : null,
          ),
          child: Padding(
            padding: const EdgeInsets.all(8.0),
            child: Row(
                mainAxisAlignment: MainAxisAlignment.start,
                children: cardContentForSize(context, widget.pr)),
          ),
        ),
      ),
    );
  }
}

List<Widget> cardContentForSize(BuildContext context, PR pr) {
  var userHash = sha1.convert(utf8.encode(pr.user));
  if (MediaQuery.of(context).orientation == Orientation.portrait) {
    return [
      Tooltip(
        message: pr.status,
        child: CircleAvatar(
          backgroundColor: Colors.transparent,
          foregroundColor: Colors.black,
          child: FaIcon(
            iconForPRStatus(pr.status),
          ),
        ),
      ),
      Expanded(
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            Flexible(
              child: Padding(
                padding: const EdgeInsets.all(8.0),
                child: Flexible(
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
                      Text(
                        DateFormat("dd MMM. yyyy").format(pr.created) +
                            " by " +
                            pr.user,
                        overflow: TextOverflow.ellipsis,
                        style: Theme.of(context).textTheme.overline,
                      ),
                    ],
                  ),
                ),
              ),
            ),
          ],
        ),
      )
    ];
  }

  return [
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
          Flexible(
            child: Padding(
              padding: const EdgeInsets.all(8.0),
              child: Flexible(
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
            ),
          ),
          Row(
            children: [
              Text(
                DateFormat("dd MMM. yyyy").format(pr.created),
                overflow: TextOverflow.ellipsis,
              ),
              Padding(
                padding: const EdgeInsets.only(left: 5),
                child: Tooltip(
                  message: pr.status,
                  child: CircleAvatar(
                    backgroundColor: Colors.transparent,
                    foregroundColor: Colors.black,
                    child: FaIcon(
                      iconForPRStatus(pr.status),
                    ),
                  ),
                ),
              ),
            ],
          ),
        ],
      ),
    ),
  ];
}

IconData iconForPRStatus(String status) {
  switch (status) {
    case "approved":
      return FontAwesomeIcons.check;
    case "closed":
      return FontAwesomeIcons.times;
    case "draft":
      return FontAwesomeIcons.firstdraft;
    default:
      return FontAwesomeIcons.codeBranch;
  }
}
