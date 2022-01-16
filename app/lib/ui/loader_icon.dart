import 'package:flutter/material.dart';

class LoaderIcon extends StatefulWidget {
  final bool isLoading;

  const LoaderIcon({Key? key, required this.isLoading}) : super(key: key);

  @override
  _LoaderIconState createState() => _LoaderIconState();
}

class _LoaderIconState extends State<LoaderIcon> with TickerProviderStateMixin {
  @override
  Widget build(BuildContext context) {
    if (!widget.isLoading) {
      return const Icon(Icons.replay_outlined);
    }

    AnimationController controller = AnimationController(
      duration: const Duration(seconds: 1),
      vsync: this,
    )..repeat(reverse: false);

    Animation<double> animation = CurvedAnimation(
      parent: controller,
      curve: Curves.ease,
    );

    return RotationTransition(
        turns: ReverseAnimation(animation),
        child: const Icon(Icons.replay_outlined));
  }
}
