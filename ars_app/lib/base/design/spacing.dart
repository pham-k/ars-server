import 'package:flutter/cupertino.dart';

class ArsSpacing {
  ArsSpacing({
    this.scale = 1,
    this.scaleVertical = 1,
    this.scaleHorizontal = 1,
  });

  final double scale;
  final double scaleVertical;
  final double scaleHorizontal;

  EdgeInsetsGeometry get screenMargin => EdgeInsets.fromLTRB(
    s(16),
    s(8),
    s(16),
    s(8),
  );

  double get inputBorderRadius => s(8);
  double get inputBorderWidth => s(2);
  EdgeInsetsGeometry get inputLabelPadding => EdgeInsets.only(
    top: s(4),
    left: s(12),
    right: s(12),
    bottom: s(8),
  );

  double get buttonMinimumHeight => s(48);
  double get buttonBorderWidth => s(2);
  double get buttonBorderRadius => s(8);

  double s(double value) {
    return value * scale;
  }

  double sv(double value) {
    return value * scaleVertical;
  }

  double sh(double value) {
    return value * scaleHorizontal;
  }
}