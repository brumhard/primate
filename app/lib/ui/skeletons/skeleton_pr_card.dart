import 'package:primate/ui/skeletons/skeleton.dart';
import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';

class PRCardSkeleton extends StatelessWidget {
  const PRCardSkeleton({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(4),
      child: AnimatedContainer(
        duration: const Duration(milliseconds: 10),
        decoration: BoxDecoration(
          boxShadow: [
            BoxShadow(color: Theme.of(context).shadowColor, blurRadius: 2)
          ],
          borderRadius: BorderRadius.circular(4),
          color: Theme.of(context).scaffoldBackgroundColor,
        ),
        child: Padding(
          padding: const EdgeInsets.all(8.0),
          child: Row(
              mainAxisAlignment: MainAxisAlignment.start,
              children: cardContentForSize(context)),
        ),
      ),
    );
  }
}

List<Widget> cardContentForSize(BuildContext context) {
  if (MediaQuery.of(context).orientation == Orientation.portrait) {
    return [
      const SizedBox(
        height: 40,
        width: 40,
        child: ClipOval(
          child: Skeleton(),
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
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Skeleton(
                      width: 200,
                      height:
                          Theme.of(context).textTheme.headline6?.fontSize ?? 20,
                    ),
                    SizedBox(
                      height:
                          Theme.of(context).textTheme.overline?.fontSize ?? 15,
                    ),
                    Skeleton(
                      width: 100,
                      height:
                          Theme.of(context).textTheme.overline?.fontSize ?? 15,
                    ),
                    SizedBox(
                      height:
                          Theme.of(context).textTheme.overline?.fontSize ?? 15,
                    ),
                    Skeleton(
                      width: 120,
                      height:
                          Theme.of(context).textTheme.overline?.fontSize ?? 15,
                    )
                  ],
                ),
              ),
            ),
          ],
        ),
      )
    ];
  }

  return [
    const SizedBox(
      height: 40,
      width: 40,
      child: ClipOval(
        child: Skeleton(),
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
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Skeleton(
                    width: 200,
                    height:
                        Theme.of(context).textTheme.headline6?.fontSize ?? 20,
                  ),
                  SizedBox(
                    height:
                        Theme.of(context).textTheme.overline?.fontSize ?? 15,
                  ),
                  Skeleton(
                    width: 120,
                    height:
                        Theme.of(context).textTheme.overline?.fontSize ?? 15,
                  ),
                ],
              ),
            ),
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
