import 'package:ars_app/base/design/design.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

enum ButtonFilledVariant {
  label,
  labelAndIcon,
}

class ButtonFilled extends StatelessWidget {
  const ButtonFilled({super.key,
    required this.label,
    this.icon,
    required this.onPressed,
    this.tooltipMessageOnDisabled,
    this.loading = false,
    this.iconAlignment = IconAlignment.start,
    this.fullWidth = false,
  }) : _variant = ButtonFilledVariant.label;

  const ButtonFilled.icon({super.key,
    required this.label,
    required this.icon,
    required this.onPressed,
    this.tooltipMessageOnDisabled,
    this.loading = false,
    this.iconAlignment = IconAlignment.start,
    this.fullWidth = false,
  }) : _variant = ButtonFilledVariant.labelAndIcon;

  final ButtonFilledVariant _variant;
  final String label;
  final Widget? icon;
  final void Function()? onPressed;
  final String? tooltipMessageOnDisabled;
  final bool loading;
  final IconAlignment iconAlignment;
  final bool fullWidth;

  @override
  Widget build(BuildContext context) {
    if (icon == null) {
      return _ButtonFilled(
        label: label,
        onPressed: onPressed,
        tooltipMessageOnDisabled: tooltipMessageOnDisabled,
        loading: loading,
        fullWidth: fullWidth,
      );
    }

    switch (_variant) {
      case ButtonFilledVariant.label:
        return _ButtonFilled(
          label: label,
          onPressed: onPressed,
          tooltipMessageOnDisabled: tooltipMessageOnDisabled,
          loading: loading,
          fullWidth: fullWidth,
        );
      case ButtonFilledVariant.labelAndIcon:
        return _ButtonFilledIcon(
          label: label,
          icon: icon!,
          onPressed: onPressed,
          tooltipMessageOnDisabled: tooltipMessageOnDisabled,
          loading: loading,
          iconAlignment: iconAlignment,
          fullWidth: fullWidth,
        );
      default:
        return _ButtonFilled(
          label: label,
          onPressed: onPressed,
          tooltipMessageOnDisabled: tooltipMessageOnDisabled,
          loading: loading,
          fullWidth: fullWidth,
        );
    }

  }
}

class _ButtonFilled extends StatelessWidget {
  const _ButtonFilled({required this.label,
    required this.onPressed,
    this.tooltipMessageOnDisabled,
    this.loading = false,
    this.fullWidth = false,
  });

  final String label;
  final void Function()? onPressed;
  final String? tooltipMessageOnDisabled;
  final bool loading;
  final bool fullWidth;

  @override
  Widget build(BuildContext context) {
    var des = Provider.of<Design>(context);
    var loc = AppLocalizations.of(context)!;

    if (loading) {
      return fullWidth ? Expanded(child: _buildButtonLoading(des, loc)) : _buildButtonLoading(des, loc);
    } else {
      return fullWidth ? Expanded(child: _buildButton(des, loc)) : _buildButton(des, loc);
    }
  }

  Widget _buildButton(Design des, AppLocalizations loc) {
    return Tooltip(
      message: _getTooltipMessageOnDisabled(loc),
      triggerMode: TooltipTriggerMode.tap,
      preferBelow: false,
      verticalOffset: des.spacing.s(32),
      child: FilledButton(
        style: _getStyle(des),
        onPressed: onPressed,
        child: Text(label),
      ),
    );
  }

  Widget _buildButtonLoading(Design des, AppLocalizations loc) {
    return FilledButton(
      style: _getStyle(des),
      onPressed: () {},
      child: const _CircularProgressIndicator(),
    );
  }

  String _getTooltipMessageOnDisabled(AppLocalizations loc) {
    if (onPressed != null) {
      return '';
    } else if (tooltipMessageOnDisabled != null) {
      return tooltipMessageOnDisabled!;
    } else {
      return loc.button_disabled;
    }
  }

  ButtonStyle? _getStyle(Design des) {
    if (fullWidth) {
      return ButtonStyle(
        minimumSize: WidgetStateProperty.all<Size?>(Size.fromHeight(des.spacing.buttonMinimumHeight)),
      );
    } else {
      return null;
    }
  }
}

class _ButtonFilledIcon extends StatelessWidget {
  const _ButtonFilledIcon({
    required this.label,
    required this.icon,
    required this.onPressed,
    this.tooltipMessageOnDisabled,
    this.loading = false,
    this.iconAlignment = IconAlignment.start,
    this.fullWidth = false,
  });

  final String label;
  final Widget icon;
  final void Function()? onPressed;
  final String? tooltipMessageOnDisabled;
  final bool loading;
  final IconAlignment iconAlignment;
  final bool fullWidth;

  @override
  Widget build(BuildContext context) {
    var des = Provider.of<Design>(context);
    var loc = AppLocalizations.of(context)!;

    if (loading) {
      return fullWidth ? Expanded(child: _buildButtonLoading(des, loc)) : _buildButtonLoading(des, loc);
    } else {
      return fullWidth ? Expanded(child: _buildButton(des, loc)) : _buildButton(des, loc);
    }
  }

  Widget _buildButton(Design des, AppLocalizations loc) {

    return Tooltip(
      message: _getTooltipMessageOnDisabled(loc),
      triggerMode: TooltipTriggerMode.tap,
      preferBelow: false,
      verticalOffset: des.spacing.s(32),
      child: FilledButton.icon(
        style: _getStyle(des),
        label: Text(label),
        icon: icon,
        onPressed: onPressed,
        iconAlignment: iconAlignment,
      ),
    );
  }

  Widget _buildButtonLoading(Design des, AppLocalizations loc) {
    return FilledButton(
      style: _getStyle(des),
      onPressed: () {},
      child: const _CircularProgressIndicator(),
    );
  }

  String _getTooltipMessageOnDisabled(AppLocalizations loc) {
    if (onPressed != null) {
      return '';
    } else if (tooltipMessageOnDisabled != null) {
      return tooltipMessageOnDisabled!;
    } else {
      return loc.button_disabled;
    }
  }

  ButtonStyle? _getStyle(Design des) {
    if (fullWidth) {
      return ButtonStyle(
        minimumSize: WidgetStateProperty.all<Size?>(Size.fromHeight(des.spacing.buttonMinimumHeight)),
      );
    } else {
      return null;
    }
  }
}

class _CircularProgressIndicator extends StatelessWidget {
  const _CircularProgressIndicator();

  @override
  Widget build(BuildContext context) {
    var des = Provider.of<Design>(context);
    return SizedBox(
      width: des.spacing.s(20),
      height: des.spacing.s(20),
      child: CircularProgressIndicator(
        strokeWidth: des.spacing.s(2),
        color: des.color.white,
      ),
    );
  }
}

