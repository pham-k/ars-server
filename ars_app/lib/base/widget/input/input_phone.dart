import 'package:ars_app/base/design/design.dart';
import 'package:flutter/material.dart';
import 'package:phone_form_field/phone_form_field.dart';
import 'package:provider/provider.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class InputPhone extends StatefulWidget {
  const InputPhone({super.key,
    required this.controller,
    this.focusNode,
    this.enabled = true,
    this.autovalidateMode = AutovalidateMode.onUserInteraction,
    this.onChanged,
    this.onSaved,
    this.validator,
    this.padding,
    this.margin,
  });

  final PhoneController controller;
  final FocusNode? focusNode;
  final bool enabled;
  final AutovalidateMode autovalidateMode;
  final void Function(PhoneNumber?)? onChanged;
  final void Function(PhoneNumber?)? onSaved;
  final String? Function(String?)? validator;
  final EdgeInsetsGeometry? padding;
  final EdgeInsetsGeometry? margin;

  @override
  State<InputPhone> createState() => _InputPhoneState();
}

class _InputPhoneState extends State<InputPhone> {
  late Design ds;
  late AppLocalizations al;

  @override
  Widget build(BuildContext context) {
    ds = Provider.of<Design>(context);
    al = AppLocalizations.of(context)!;

      return Container(
        padding: widget.padding,
        margin: widget.margin ?? EdgeInsets.only(bottom: ds.spacing.s(4)),
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
    TextStyle style;

    if (!widget.enabled) {
      style = ds.typo.inputLabel.copyWith(color: ds.color.grey);
    } else {
      style = ds.typo.inputLabel;
    }
    return Text(
      al.input_phone_label,
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
    return PhoneFormField(
      controller: widget.controller,
      focusNode: widget.focusNode,
      enabled: widget.enabled,
      countrySelectorNavigator: const CountrySelectorNavigator.modalBottomSheet(
          favorites: [IsoCode.VN]
      ),
      autovalidateMode: widget.autovalidateMode,
      decoration: _getFieldDecoration(),
      validator: (PhoneNumber? value) {
        return PhoneValidator.compose([
          PhoneValidator.required(context, errorText: "aaa"),
          PhoneValidator.validMobile(context, errorText: "aaa",)
        ])(value);
      },
      onChanged: widget.onChanged,
      onSaved: widget.onSaved,
    );
  }

  Widget _getSuffixIcon() {
    if (!widget.enabled) {
      return IconButton(
        onPressed: () {},
        icon: const Icon(Icons.block_outlined),
      );
    } else {
      return IconButton(
        onPressed:() {
          widget.controller.changeNationalNumber('');
        },
        icon: const Icon(Icons.clear),
      );
    }
  }

  InputDecoration _getFieldDecoration() {
    return InputDecoration(
      hintText: al.input_phone_hint,
      helperText: al.input_phone_helper,
      suffixIcon: _getSuffixIcon(),
    );
  }
}