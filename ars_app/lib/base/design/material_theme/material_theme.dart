import "package:ars_app/base/design/design.dart";
import "package:ars_app/base/design/material_theme/button_theme.dart";
import "package:ars_app/base/design/material_theme/card_theme.dart";
import "package:ars_app/base/design/material_theme/checkbox_theme.dart";
import "package:ars_app/base/design/material_theme/color_scheme.dart";
import "package:ars_app/base/design/material_theme/input_decoration.dart";
import "package:ars_app/base/design/material_theme/text_theme.dart";
import "package:ars_app/base/design/material_theme/tooltip_theme.dart";
import "package:flutter/material.dart";

class MaterialTheme {
  final Design des;

  const MaterialTheme(
      this.des
  );

  ThemeData get themeData => ThemeData(
    useMaterial3: true,
    brightness: des.brightness,
    colorScheme: getColorScheme(des),
    fontFamily: "HelveticaNeue",
    textTheme: getTextTheme(des),
    scaffoldBackgroundColor: des.color.surface,
    canvasColor: des.color.surface,
    filledButtonTheme: getFilledButtonTheme(des),
    outlinedButtonTheme: getOutlinedButtonTheme(des),
    tooltipTheme: getTooltipTheme(des),
    inputDecorationTheme: getInputDecorationTheme(des),
    checkboxTheme: getCheckBoxTheme(des),
    cardTheme: getCardTheme(des),
  );
}
