import 'package:ars_app/base/design/color.dart';
import 'package:flutter/material.dart';

class ArsTypography {
  ArsTypography({
    required this.color
  });

  final ArsColor color;

  static const String _fontFamily = 'HelveticaNeue';

  TextStyle _getTextStyle({
    required FontWeight fontWeight,
    required double fontSize,
    Color? color,
  }) {
    return TextStyle(
      fontSize: fontSize,
      fontWeight: fontWeight,
      fontFamily: _fontFamily,
      height: 1,
      color: color,
    );
  }

  // Body
  TextStyle get bodyNormal => _getTextStyle(
    fontWeight: FontWeight.w500,
    fontSize: 16,
    color: color.black,
  );

  TextStyle get bodyMedium => _getTextStyle(
    fontWeight: FontWeight.w700,
    fontSize: 16,
    color: color.black,
  );

  TextStyle get bodyBold => _getTextStyle(
    fontWeight: FontWeight.w900,
    fontSize: 16,
    color: color.black,
  );

  // Heading 1
  TextStyle get h1Normal => _getTextStyle(
    fontWeight: FontWeight.w500,
    fontSize: 24,
    color: color.black,
  );

  TextStyle get h1Medium => _getTextStyle(
    fontWeight: FontWeight.w700,
    fontSize: 24,
    color: color.black,
  );

  TextStyle get h1Bold => _getTextStyle(
    fontWeight: FontWeight.w900,
    fontSize: 24,
    color: color.black,
  );

  // Heading 2
  TextStyle get h2Normal => _getTextStyle(
    fontWeight: FontWeight.w500,
    fontSize: 20,
    color: color.black,
  );

  TextStyle get h2Medium => _getTextStyle(
    fontWeight: FontWeight.w700,
    fontSize: 20,
    color: color.black,
  );

  TextStyle get h2Bold => _getTextStyle(
    fontWeight: FontWeight.w900,
    fontSize: 20,
    color: color.black,
  );

  // Button label
  TextStyle get buttonLabel => _getTextStyle(
    fontWeight: FontWeight.w700,
    fontSize: 16,
    color: color.black,
  );

  TextStyle get inputLabel => _getTextStyle(
    fontWeight: FontWeight.w700,
    fontSize: 18,
    color: color.black,
  );

  TextStyle get inputHint => _getTextStyle(
    fontWeight: FontWeight.w500,
    fontSize: 16,
    color: color.grey,
  );

  TextStyle get inputHelper => _getTextStyle(
    fontWeight: FontWeight.w500,
    fontSize: 14,
    color: color.grey,
  );

  TextStyle get inputError => _getTextStyle(
    fontWeight: FontWeight.w500,
    fontSize: 14,
    color: color.red,
  );

  TextStyle get tooltip => _getTextStyle(
    fontWeight: FontWeight.w500,
    fontSize: 16,
    color: color.white,
  );
}
