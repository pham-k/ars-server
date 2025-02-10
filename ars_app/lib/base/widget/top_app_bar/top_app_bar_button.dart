import 'package:ars_app/base/design/design.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

enum TopAppBarIcon {
  menu,
  chevronLeft,
  moreVertical,
  settings
}

class TopAppBarButton extends StatelessWidget {
  const TopAppBarButton({super.key,
    required this.icon,
    this.onTap,
  });

  final TopAppBarIcon icon;
  final void Function()? onTap;

  @override
  Widget build(BuildContext context) {
    Design des = Provider.of<Design>(context);
    double iconSize = 24;
    double buttonSize = 40;

    return InkWell(
      onTap: onTap,
      borderRadius: BorderRadius.circular(des.spacing.s(12)),
      child: Center(
        widthFactor: buttonSize / iconSize,
        heightFactor: buttonSize / iconSize,
        child: _buildIcon(des),
      ),
    );
  }

  Widget _buildIcon(Design ds) {
    switch (icon) {
      case TopAppBarIcon.settings:
        return Icon(
          Icons.settings_outlined,
          color: ds.color.black,
          size: 24,
        );
      case TopAppBarIcon.menu:
        return Icon(
          Icons.menu,
          color: ds.color.black,
          size: 24,
        );
      case TopAppBarIcon.chevronLeft:
        return Icon(
          Icons.chevron_left,
          color: ds.color.black,
          size: 32,
        );
      default:
        return Icon(
          Icons.person_outlined,
          color: ds.color.black,
          size: 24,
        );
    }
  }
}
