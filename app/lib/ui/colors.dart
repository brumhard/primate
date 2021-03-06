import 'package:flutter/material.dart';

const Color pink = Color(0xFFdd5e89);
const Color orange = Color(0xFFf7bb97);

MaterialColor createMaterialColor(Color color) {
  List strengths = <double>[.05];
  final swatch = <int, Color>{};
  final int r = color.red, g = color.green, b = color.blue;

  for (int i = 1; i < 10; i++) {
    strengths.add(0.1 * i);
  }
  for (var strength in strengths) {
    final double ds = 0.5 - strength;
    swatch[(strength * 1000).round()] = Color.fromRGBO(
      r + ((ds < 0 ? r : (255 - r)) * ds).round(),
      g + ((ds < 0 ? g : (255 - g)) * ds).round(),
      b + ((ds < 0 ? b : (255 - b)) * ds).round(),
      1,
    );
  }
  return MaterialColor(color.value, swatch);
}

extension CustomColorScheme on ColorScheme {
  Color get fresh => brightness == Brightness.light
      ? const Color(0xFFc9d3f8)
      : const Color(0xFF3b60e4);
  Color get waiting => brightness == Brightness.light
      ? const Color(0xFFd7eae0)
      : const Color(0xFF519872);
  Color get stale => brightness == Brightness.light
      ? const Color(0xFFf4b8c2)
      : const Color(0xFF721121);
  Color get rotten => brightness == Brightness.light
      ? const Color(0xFFebdbc1)
      : const Color(0xFF563f1b);
}
