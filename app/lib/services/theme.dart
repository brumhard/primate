import 'package:primate/ui/colors.dart';
import 'package:flutter/material.dart';

// TODO: save to shared preferences
class ThemeService extends ChangeNotifier {
  final Map<Brightness, ThemeData> _themes = {
    Brightness.dark: ThemeData(
      brightness: Brightness.dark,
      primarySwatch: createMaterialColor(orange),
      iconTheme: const IconThemeData(color: Colors.white),
    ),
    Brightness.light: ThemeData(
      brightness: Brightness.light,
      primarySwatch: createMaterialColor(pink),
      iconTheme: const IconThemeData(color: Colors.black),
    ),
  };

  Brightness _currentBrightness = Brightness.dark;

  ThemeService();

  ThemeData get theme => _themes[_currentBrightness]!;
  Brightness get currentBrightness => _currentBrightness;

  void setTheme(Brightness brightness) {
    _currentBrightness = brightness;
    notifyListeners();
  }
}
