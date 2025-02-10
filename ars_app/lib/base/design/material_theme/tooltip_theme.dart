import 'package:ars_app/base/design/design.dart';
import 'package:flutter/material.dart';

TooltipThemeData getTooltipTheme(Design ds) {
  return TooltipThemeData(
      decoration: ds.decor.tooltip,
      textStyle: ds.typo.tooltip
  );
}