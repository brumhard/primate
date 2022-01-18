import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';

class TurnAtLeastOnceOnListenable extends StatefulWidget {
  final ValueListenable<bool> listenable;
  final Widget child;

  const TurnAtLeastOnceOnListenable(
      {Key? key, required this.listenable, required this.child})
      : super(key: key);

  @override
  _TurnAtLeastOnceOnListenableState createState() =>
      _TurnAtLeastOnceOnListenableState();
}

class _TurnAtLeastOnceOnListenableState
    extends State<TurnAtLeastOnceOnListenable> with TickerProviderStateMixin {
  bool _shoodRotate = false;
  late AnimationController _controller;
  late Animation<double> _animation;

  @override
  void initState() {
    _controller = AnimationController(
      duration: const Duration(seconds: 1),
      vsync: this,
    )..addStatusListener((status) {
        if (status == AnimationStatus.completed && _shoodRotate) {
          _controller.forward(from: 0.0);
        }
      });

    widget.listenable.addListener(() {
      _shoodRotate = widget.listenable.value;
      if (_shoodRotate) {
        // trigger first animation cycle
        _controller.forward(from: 0.0);
      }
    });

    _animation = ReverseAnimation(CurvedAnimation(
      parent: _controller,
      curve: Curves.ease,
    ));

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return RotationTransition(
      turns: _animation,
      child: widget.child,
    );
  }
}
