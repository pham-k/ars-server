import 'package:ars_app/base/design/design.dart';
import 'package:flutter/material.dart';

InputDecorationTheme getInputDecorationTheme(Design ds) {
  return InputDecorationTheme(
    // Label
    floatingLabelBehavior: FloatingLabelBehavior.never,
    floatingLabelAlignment: FloatingLabelAlignment.start,
    // Hint
    hintStyle: ds.typo.inputHint,
    // Helper
    helperStyle: ds.typo.inputHelper,
    errorStyle: ds.typo.inputError,
    errorMaxLines: 1,
    // Icons
    suffixIconColor: WidgetStateColor.resolveWith((Set<WidgetState> states) {
      if (states.contains(WidgetState.disabled) &&
          !states.contains(WidgetState.focused)) {
        return ds.color.grey;
      }
      if (states.contains(WidgetState.error)) {
        return ds.color.red;
      }
      if (states.contains(WidgetState.focused)) {
        return ds.color.blue;
      }
      return ds.color.black;
    }),
    // Border
    border: OutlineInputBorder(
      borderSide: BorderSide(width: ds.spacing.s(2), color: ds.color.black),
      borderRadius: BorderRadius.circular(ds.spacing.s(8)),
    ),
    enabledBorder: OutlineInputBorder(
      borderSide: BorderSide(width: ds.spacing.s(2), color: ds.color.black),
      borderRadius: BorderRadius.circular(ds.spacing.s(8)),
    ),
    focusedBorder: OutlineInputBorder(
      borderSide: BorderSide(width: ds.spacing.s(2), color: ds.color.blue),
      borderRadius: BorderRadius.circular(ds.spacing.s(8)),
    ),
    errorBorder: OutlineInputBorder(
      borderSide: BorderSide(width: ds.spacing.s(2), color: ds.color.red),
      borderRadius: BorderRadius.circular(ds.spacing.s(8)),
    ),
    focusedErrorBorder: OutlineInputBorder(
      borderSide: BorderSide(width: ds.spacing.s(2), color: ds.color.red),
      borderRadius: BorderRadius.circular(ds.spacing.s(8)),
    ),
    disabledBorder: OutlineInputBorder(
      borderSide: BorderSide(width: ds.spacing.s(2), color: ds.color.grey),
      borderRadius: BorderRadius.circular(ds.spacing.s(8)),
    ),
  );
}