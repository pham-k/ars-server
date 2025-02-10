import 'dart:math';

import 'package:ars_app/base/design/design.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter_animate/flutter_animate.dart';
import 'package:provider/provider.dart';

class ArsLogo extends StatelessWidget {
  const ArsLogo({super.key,
    this.size = 24,
    this.animate = false,
  });

  final double size;
  final bool animate;

  @override
  Widget build(BuildContext context) {
    var des = Provider.of<Design>(context);

    if (animate) {
      return ClipPath(
        clipper: _TriangleClipper(),
        child: Container(
          color: des.color.blue,
          height: size * sqrt(3) / 2,
          width: size,
        ),
      ).animate(
        // delay: 500.ms,
        onPlay: (controller) => controller.repeat(reverse: true),
      ).fadeIn(duration: 800.ms, delay: 500.ms, curve: Curves.easeIn);
    } else {
      return ClipPath(
        clipper: _TriangleClipper(),
        child: Container(
          color: des.color.blue,
          height: size * sqrt(3) / 2,
          width: size,
        ),
      );
    }

  }
}

class _TriangleClipper extends CustomClipper<Path> {
  @override
  Path getClip(Size size) {
    final path = Path();
    path.lineTo(size.width, 0.0);
    path.lineTo(size.width / 2, size.height);
    path.close();
    return path;
  }

  @override
  bool shouldReclip(_TriangleClipper oldClipper) => false;
}
