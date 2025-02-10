import 'package:ars_app/base/design/design.dart';
import 'package:flutter/material.dart';

FilledButtonThemeData getFilledButtonTheme(Design ds) {
  return const FilledButtonThemeData();
}

OutlinedButtonThemeData getOutlinedButtonTheme(Design ds) {
  return OutlinedButtonThemeData(
    style: ButtonStyle(
      // minimumSize: WidgetStateProperty.all<Size?>(Size.fromHeight(ds.sp.s(48))),
      side: _getOutlinedButtonSide(ds),
    ),
  );
}

WidgetStateProperty<BorderSide?> _getOutlinedButtonSide(Design ds) {
  return WidgetStateProperty.resolveWith<BorderSide?>((Set<WidgetState> states) {
    if (states.contains(WidgetState.disabled)) {
      return BorderSide(
          color: ds.color.grey,
          width: ds.spacing.buttonBorderWidth
      );
    } else {
      return BorderSide(
          color: ds.color.blue,
          width: ds.spacing.buttonBorderWidth
      );
    }
  },);
}
