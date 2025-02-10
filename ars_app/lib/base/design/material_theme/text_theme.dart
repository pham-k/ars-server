import 'package:ars_app/base/design/design.dart';
import 'package:flutter/material.dart';

TextTheme getTextTheme(Design ds) {
  return const TextTheme(
    titleLarge: TextStyle(
      fontWeight: FontWeight.w500,
      fontSize: 24,
    ),
    titleMedium: TextStyle(
      fontWeight: FontWeight.w500,
      fontSize: 22,
    ),
    titleSmall: TextStyle(
      fontWeight: FontWeight.w500,
      fontSize: 20,
    ),
    bodyLarge: TextStyle(
      fontWeight: FontWeight.w500,
      fontSize: 20,
    ),
    bodyMedium: TextStyle(
      fontWeight: FontWeight.w500,
      fontSize: 18,
    ),
    bodySmall: TextStyle(
      fontWeight: FontWeight.w500,
      fontSize: 16,
    ),
    labelLarge: TextStyle(
      fontWeight: FontWeight.w700,
      fontSize: 18,
    ),
    labelMedium: TextStyle(
      fontWeight: FontWeight.w700,
      fontSize: 16,
    ),
    labelSmall: TextStyle(
      fontWeight: FontWeight.w700,
      fontSize: 14,
    ),
  ).apply(
    bodyColor: ds.color.black,
    displayColor: ds.color.black,
  );
}