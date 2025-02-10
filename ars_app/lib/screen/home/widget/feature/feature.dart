import 'package:ars_app/base/design/design.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:provider/provider.dart';

class Feature extends StatelessWidget {
  const Feature({super.key,
    required this.name,
    required this.image,
    this.onTap,
    this.margin,
    this.padding,
  });

  final String name;
  final String image;
  final void Function()? onTap;
  final EdgeInsetsGeometry? margin;
  final EdgeInsetsGeometry? padding;

  @override
  Widget build(BuildContext context) {
    Design des = Provider.of<Design>(context);

    return Container(
      margin: margin ?? EdgeInsets.only(bottom: des.spacing.s(16)),
      padding: padding,
      child: InkWell(
        onTap: onTap,
        borderRadius: BorderRadius.circular(des.spacing.s(12)),
        child: Card.filled(
          margin: EdgeInsets.zero,
          child: Stack(
            children: [
              Positioned(
                right: des.spacing.s(16),
                bottom: 0,
                child: ClipRRect(
                  borderRadius: const BorderRadius.only(
                    bottomRight: Radius.circular(12),
                  ),
                  child: SvgPicture.asset(
                    image,
                    width: des.spacing.s(120),
                  ),
                ),
              ),
              Container(
                width: double.infinity,
                padding: EdgeInsets.symmetric(
                    horizontal: des.spacing.s(16),
                    vertical: des.spacing.s(40)
                ),
                child: Text(name, style: des.typo.h1Medium),
              )
            ],
          ),
        ),
      ),
    );
  }
}
