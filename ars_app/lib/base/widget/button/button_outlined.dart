import 'package:ars_app/base/design/design.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

enum ButtonOutlinedVariant {
  label,
  labelAndIcon,
}

class ButtonOutlined extends StatelessWidget {
  const ButtonOutlined({super.key,
    required this.label,
    this.icon,
    required this.onPressed,
    this.tooltipMessageOnDisabled,
    this.loading = false,
    this.iconAlignment = IconAlignment.start,
    this.fullWidth = false,
  }) : _variant = ButtonOutlinedVariant.label;

  const ButtonOutlined.icon({super.key,
    required this.label,
    required this.icon,
    required this.onPressed,
    this.tooltipMessageOnDisabled,
    this.loading = false,
    this.iconAlignment = IconAlignment.start,
    this.fullWidth = false,
  }) : _variant = ButtonOutlinedVariant.labelAndIcon;

  final ButtonOutlinedVariant _variant;
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
      return _ButtonOutlined(
        label: label,
        onPressed: onPressed,
        tooltipMessageOnDisabled: tooltipMessageOnDisabled,
        loading: loading,
        fullWidth: fullWidth,
      );
    }

    switch (_variant) {
      case ButtonOutlinedVariant.label:
        return _ButtonOutlined(
          label: label,
          onPressed: onPressed,
          tooltipMessageOnDisabled: tooltipMessageOnDisabled,
          loading: loading,
          fullWidth: fullWidth,
        );
      case ButtonOutlinedVariant.labelAndIcon:
        return _ButtonOutlinedIcon(
          label: label,
          icon: icon!,
          onPressed: onPressed,
          tooltipMessageOnDisabled: tooltipMessageOnDisabled,
          loading: loading,
          iconAlignment: iconAlignment,
          fullWidth: fullWidth,
        );
      default:
        return _ButtonOutlined(
          label: label,
          onPressed: onPressed,
          tooltipMessageOnDisabled: tooltipMessageOnDisabled,
          loading: loading,
          fullWidth: fullWidth,
        );
    }

  }
}

class _ButtonOutlined extends StatelessWidget {
  const _ButtonOutlined({required this.label,
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
      child: OutlinedButton(
        style: _getStyle(des),
        onPressed: onPressed,
        child: Text(label),
      ),
    );
  }

  Widget _buildButtonLoading(Design des, AppLocalizations loc) {
    return OutlinedButton(
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

class _ButtonOutlinedIcon extends StatelessWidget {
  const _ButtonOutlinedIcon({
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
      child: OutlinedButton.icon(
        style: _getStyle(des),
        label: Text(label),
        icon: icon,
        onPressed: onPressed,
        iconAlignment: iconAlignment,
      ),
    );
  }

  Widget _buildButtonLoading(Design des, AppLocalizations loc) {
    return OutlinedButton(
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
  const _CircularProgressIndicator({super.key});

  @override
  Widget build(BuildContext context) {
    var ds = Provider.of<Design>(context);
    return SizedBox(
      width: ds.spacing.s(20),
      height: ds.spacing.s(20),
      child: CircularProgressIndicator(
        strokeWidth: ds.spacing.s(2),
        color: ds.color.blue,
      ),
    );
  }
}

