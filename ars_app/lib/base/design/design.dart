import 'package:ars_app/base/design/decoration.dart';
import 'package:ars_app/base/design/color.dart';
import 'package:ars_app/base/design/spacing.dart';
import 'package:ars_app/base/design/typography.dart';
import 'package:flutter/material.dart';
import 'package:rxdart/rxdart.dart';

enum ArsTheme {
  light,
  dark,
}

class Design {
  Design({
    this.theme = ArsTheme.light,
    this.themeMode = ThemeMode.light,
    required this.onThemeChanged,
  });

  // Material data
  ThemeMode themeMode;
  Brightness get brightness {
    return Brightness.light;
  }

  PublishSubject<ArsTheme> onThemeChanged;

  ArsTheme theme;
  ArsColor get color {
    return ArsColor.light();
  }

  ArsTypography get typo => ArsTypography(color: color);
  ArsSpacing get spacing => ArsSpacing();
  ArsDecoration get decor => ArsDecoration(color: color, spacing: spacing, typo: typo);

  void setTheme(ArsTheme newTheme) {
    theme = newTheme;
    onThemeChanged.add(newTheme);
  }
}