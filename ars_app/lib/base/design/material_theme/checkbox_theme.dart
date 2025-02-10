import 'package:flutter/material.dart';

import '../design.dart';

CheckboxThemeData getCheckBoxTheme(Design des) {
  return const CheckboxThemeData(
    visualDensity: VisualDensity.compact,
    materialTapTargetSize: MaterialTapTargetSize.shrinkWrap,
  );
}