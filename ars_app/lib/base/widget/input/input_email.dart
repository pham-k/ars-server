import 'package:ars_app/base/design/design.dart';
import 'package:email_validator/email_validator.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class InputEmail extends StatefulWidget {
  const InputEmail({super.key,
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
  State<InputEmail> createState() => _InputEmailState();
}

class _InputEmailState extends State<InputEmail> {
  late Design ds;
  late AppLocalizations al;
  final _controller = WidgetStatesController();

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
            verticalOffset: ds.spacing.s(48),
            child: _buildField()
          ),
        ],
      ),
    );
  }

  Widget _buildLabel() {
    TextStyle? style;

    if (!widget.enabled) {
      style = ds.typo.inputLabel.copyWith(color: ds.color.grey);
    } else {
      style = ds.typo.inputLabel;
    }
    return Text(
      widget.labelText ?? al.input_email_label,
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
      autovalidateMode: widget.autovalidateMode,
      decoration: _getFieldDecoration(),
      validator: (String? value) {
        bool isValid = EmailValidator.validate(value!);
        if (!isValid) {
          return al.input_email_error_invalid;
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
        icon: const Icon(Icons.block_outlined),
      );
    } else {
      return IconButton(
        onPressed: widget.controller.clear,
        icon: const Icon(Icons.clear),
      );
    }
  }

  InputDecoration _getFieldDecoration() {
    return InputDecoration(
      hintText: widget.hintText ?? al.input_email_hint,
      helperText: widget.helperText ?? al.input_email_helper,
      suffixIcon: _getSuffixIcon(),
    );
  }
}
