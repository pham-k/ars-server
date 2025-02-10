import 'package:ars_app/base/design/design.dart';
import 'package:ars_app/base/widget/ars_logo/ars_logo.dart';
import 'package:ars_app/screen/home/home_screen.dart';
import 'package:ars_app/screen/log_in_method/log_in_method_screen.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class SplashScreen extends StatefulWidget {
  const SplashScreen({super.key});

  static const routeName = '/';

  @override
  State<SplashScreen> createState() => _SplashScreenState();
}

class _SplashScreenState extends State<SplashScreen> {
  late Design _des;
  late AppLocalizations _loc;

  bool _isFirstRun = true;

  @override
  void initState() {
    super.initState();
  }

  @override
  void dispose() {
    super.dispose();
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();

    if (_isFirstRun) {
      _onFirstRun().then((_) {
        _isFirstRun = false;
        if (mounted) {
          Navigator.of(context).pushNamedAndRemoveUntil(HomeScreen.routeName, (_) => false);
        }
      });
    }
  }

  Future<void> _onFirstRun() async {
    await Future.delayed(const Duration(seconds: 5));
  }

  @override
  Widget build(BuildContext context) {
    _des = Provider.of<Design>(context);
    _loc = AppLocalizations.of(context)!;

    return Scaffold(
      body: Center(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.center,
          mainAxisSize: MainAxisSize.max,
          children: [
            const Spacer(),
            ArsLogo(
              size: _des.spacing.s(64),
              animate: true,
            ),
            SizedBox(height: _des.spacing.s(16),),
            Text('ARS', style: _des.typo.h1Bold,),
            const Spacer(),
          ],
        ),
      ),
    );
  }
}
