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
  Color get fresh => brightness == Brightness.light? const Color.fromARGB(255, 201, 211, 248) : const Color.fromARGB(255, 59, 96, 228);
  Color get waiting => brightness == Brightness.light? const Color.fromARGB(255, 215, 234, 224) : const Color.fromARGB(255, 81, 152, 114);
  Color get stale => brightness == Brightness.light? const Color.fromARGB(255, 244, 184, 194) : const Color.fromARGB(255, 114, 17, 33);
  Color get rotten => brightness == Brightness.light? const Color.fromARGB(255, 235, 219, 193) : const Color.fromARGB(255, 86, 63, 27);
}
