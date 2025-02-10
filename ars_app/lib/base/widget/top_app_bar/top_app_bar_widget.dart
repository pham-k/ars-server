import 'package:ars_app/base/design/design.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class TopAppBar extends StatelessWidget {
  const TopAppBar({
    super.key,
    this.title = '',
    this.leading,
    this.actions
  });

  final String title;
  final Widget? leading;
  final List<Widget>? actions;


  @override
  Widget build(BuildContext context) {
    Design des = Provider.of<Design>(context);
    double paddingHorizontal = 12;
    double height = 40;
    double width = double.infinity;

    return AppBar(
      toolbarHeight: height,
      leadingWidth: height + paddingHorizontal,
      leading: Padding(
        padding: EdgeInsets.only(left: paddingHorizontal),
        child: leading ?? Container(),
      ),
      title: Text(
        title,
        style: des.typo.bodyBold,
      ),
      centerTitle: true,
      actions: [
        ...?actions,
        SizedBox(
          width: paddingHorizontal,
        )
      ],
    );
  }
}
