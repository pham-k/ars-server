import 'package:ars_app/base/design/design.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class InputPassword extends StatefulWidget {
  const InputPassword({super.key,
    required this.controller,
    this.focusNode,
    this.enabled = true,
    this.readOnly = false,
    this.autovalidateMode = AutovalidateMode.onUserInteraction,
    this.onChanged,
    this.onSaved,
    this.validator,
    this.labelText,
    this.hintText,
    this.helperText,
    this.padding,
    this.margin,
  });

  final TextEditingController controller;
  final FocusNode? focusNode;
  final bool enabled;
  final bool readOnly;
  final AutovalidateMode autovalidateMode;
  final void Function(String?)? onChanged;
  final void Function(String?)? onSaved;
  final String? Function(String?)? validator;
  final String? labelText;
  final String? hintText;
  final String? helperText;
  final EdgeInsetsGeometry? padding;
  final EdgeInsetsGeometry? margin;

  @override
  State<InputPassword> createState() => _InputPasswordState();
}

class _InputPasswordState extends State<InputPassword> {
  late Design ds;
  late AppLocalizations al;

  bool _obscureText = true;

  @override
  Widget build(BuildContext context) {
    ds = Provider.of<Design>(context);
    al = AppLocalizations.of(context)!;

      return Container(
        padding: widget.padding,
        margin: widget.margin ?? EdgeInsets.only(bottom: ds.spacing.s(8)),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Container(
              padding: ds.spacing.inputLabelPadding,
              child: Row(
                crossAxisAlignment: CrossAxisAlignment.center,
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  _buildLabel(),
                  !widget.enabled ? _buildChipDisabled() : Container(),
                ],
              ),
            ),
            Tooltip(
                message: !widget.enabled ? al.input_disabled : "",
                decoration: ds.decor.tooltip,
                textStyle: ds.typo.tooltip,
                triggerMode: TooltipTriggerMode.tap,
                preferBelow: false,
                verticalOffset: ds.spacing.s(40),
                child: _buildField()
            ),
          ],
        ),
      );
  }

  Widget _buildLabel() {
    TextStyle style;

    if (!widget.enabled) {
      style = ds.typo.inputLabel.copyWith(color: ds.color.grey);
    } else {
      style = ds.typo.inputLabel;
    }
    return Text(
      widget.labelText ?? al.input_password_label,
      style: style,
    );
  }

  Widget _buildChipDisabled() {
    return Container(
      padding: EdgeInsets.fromLTRB(
        ds.spacing.s(8),
        ds.spacing.s(4),
        ds.spacing.s(8),
        ds.spacing.s(4),
      ),
      decoration: ds.decor.chipDisabled,
      child: Text(al.input_disabled, style: ds.typo.tooltip.copyWith(color: ds.color.grey),),
    );
  }

  Widget _buildField() {
    return TextFormField(
      controller: widget.controller,
      focusNode: widget.focusNode,
      enabled: widget.enabled,
      readOnly: widget.readOnly,
      obscureText: _obscureText,
      autovalidateMode: widget.autovalidateMode,
      decoration: _getFieldDecoration(),
      validator: (String? value) {
        if (value == null || value.length < 8) {
          return al.input_password_error_too_short;
        } else if (value.length > 128) {
          return al.input_password_error_too_long;
        } else {
          return null;
        }
      },
      onChanged: widget.onChanged,
      onSaved: widget.onSaved,
    );
  }

  Widget? _getSuffixIcon() {
    if (widget.readOnly) {
      return null;
    }
    else if (!widget.enabled) {
      return IconButton(
        onPressed: () {},
        icon: const Icon(Icons.visibility_outlined),
      );
    } else if (_obscureText) {
      return IconButton(
        onPressed: () {
          _obscureText = !_obscureText;
          setState(() {});
        },
        icon: const Icon(Icons.visibility_outlined),
      );
    } else {
      return IconButton(
        onPressed: () {
          _obscureText = !_obscureText;
          setState(() {});
        },
        icon: const Icon(Icons.visibility_off_outlined),
      );
    }
  }

  InputDecoration _getFieldDecoration() {
    return InputDecoration(
      hintText: widget.hintText ?? al.input_password_hint,
      helperText: widget.helperText ?? al.input_password_helper,
      suffixIcon: _getSuffixIcon(),
    );
  }

}