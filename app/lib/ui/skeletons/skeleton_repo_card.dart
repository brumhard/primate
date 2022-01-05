import 'package:app/ui/skeletons/skeleton.dart';
import 'package:app/ui/skeletons/skeleton_pr_card.dart';
import 'package:flutter/material.dart';

class RepositoryCardSkeleton extends StatelessWidget {
  const RepositoryCardSkeleton({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Card(
      color: const Color(0xffedf2f7),
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
                    child: Skeleton(
                      width: 150,
                      height:
                          Theme.of(context).textTheme.headline5?.fontSize ?? 20,
                    ),
                  ),
                ],
              ),
            ),
            Row(
              children: [
                Flexible(flex: 1, child: Container()),
                Flexible(
                    flex: 5,
                    child: Column(
                      children: const [
                        PRCardSkeleton(),
                        PRCardSkeleton(),
                      ],
                    )),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
