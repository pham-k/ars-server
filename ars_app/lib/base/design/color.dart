import 'dart:ui';

class Palette {
  static const black = Color(0xFF161616);
  static const white = Color(0xFFF8F8F8);

  static const blue = Color(0xFF004488);
  static const red = Color(0xFFBB5566);
  static const green = Color(0xFF44AA99);
  static const yellow = Color(0xFFEECC66);
  static const grey = Color(0xFF808080);
  static const surface = Color(0xfff8f0f0);
  static const container = Color(0xffe8e0e0);

  static const blueD = Color(0xFFaaccff);
  static const redD = Color(0xFFffaabb);
  static const greenD = Color(0xFF77ddcc);
  static const yellowD = Color(0xFFffeebb);
  static const greyD = Color(0xFFeef0f1);
  static const surfaceD = Color(0xff1c1010);
  static const containerD = Color(0xff383030);
}

class ArsColor {
  ArsColor({
    required this.black,
    required this.blue,
    required this.red,
    required this.green,
    required this.yellow,
    required this.white,
    required this.grey,
    required this.surface,
    required this.container,
  });

  ArsColor.light({
    this.black = Palette.black,
    this.blue = Palette.blue,
    this.red = Palette.red,
    this.green = Palette.green,
    this.yellow = Palette.yellow,
    this.white = Palette.white,
    this.grey = Palette.grey,
    this.surface = Palette.surface,
    this.container = Palette.container,
  });

  final Color black;
  final Color blue;
  final Color red;
  final Color green;
  final Color yellow;
  final Color white;
  final Color grey;
  final Color surface;
  final Color container;

  double get overlayOpacity => 0.38;
}