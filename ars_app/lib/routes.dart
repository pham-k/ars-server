import 'package:ars_app/screen/home/home_screen.dart';
import 'package:ars_app/screen/log_in_method/log_in_method_screen.dart';
import 'package:ars_app/screen/menu/menu_screen.dart';
import 'package:ars_app/screen/splash/splash_screen.dart';
import 'package:flutter/material.dart';

class Routes {
  Routes() {
    _routes = {
      SplashScreen.routeName: (_) => const SplashScreen(),
      LogInMethodScreen.routeName: (_) => const LogInMethodScreen(),
      HomeScreen.routeName: (_) => const HomeScreen(),
      MenuScreen.routeName: (_) => const MenuScreen(),
    };
  }


  late Map<String, WidgetBuilder> _routes;

  Map<String, WidgetBuilder> get routes => _routes;
}
