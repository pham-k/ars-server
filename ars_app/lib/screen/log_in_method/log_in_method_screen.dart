import 'package:ars_app/base/design/design.dart';
import 'package:ars_app/base/widget/button/button.dart';
import 'package:ars_app/base/widget/button/button_filled.dart';
import 'package:ars_app/base/widget/developer/developer.dart';
import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:provider/provider.dart';

class LogInMethodScreen extends StatefulWidget {
  const LogInMethodScreen({super.key});

  static const routeName = '/log-in-method';

  @override
  State<LogInMethodScreen> createState() => _LogInMethodScreenState();
}

class _LogInMethodScreenState extends State<LogInMethodScreen> {
  late Design _ds;
  late AppLocalizations _al;

  bool _logInAnonymouslyIsLoading = false;
  bool _logInEmailIsLoading = false;

  @override
  Widget build(BuildContext context) {
    _ds = Provider.of<Design>(context);
    _al = AppLocalizations.of(context)!;

    Widget body = _buildBody();

    return _buildLayout(body);
  }

  Widget _buildLayout(Widget body) {
    return PopScope(
      canPop: false,
      onPopInvokedWithResult: _onPopInvokedWithResult,
      child: RefreshIndicator(
        onRefresh: _onRefresh,
        child: Scaffold(
          // backgroundColor: _ds.pl.white,
          body: SafeArea(
            child: CustomScrollView(
              slivers: [
                SliverFillRemaining(
                  hasScrollBody: true,
                  child: Container(
                    margin: _ds.spacing.screenMargin,
                    child: body
                  ),
                )
              ],
            ),
          ),
        ),
      ),
    );
  }

  Widget _buildBody() {
    return const Dev();
  }

  void _onPopInvokedWithResult (bool didPop, result) {

  }

  Future<void> _logInAnonymously() async {
    _logInAnonymouslyIsLoading = true;
    setState(() {});

    await Future.delayed(const Duration(seconds: 2));

    _logInAnonymouslyIsLoading = false;
    setState(() {});

    // if (!mounted) return;
    // Navigator.of(context).pushNamedAndRemoveUntil(HomeScreen.routeName, (_) => false);
  }

  Future<void> _logInEmail() async {
    _logInEmailIsLoading = true;
    setState(() {});

    await Future.delayed(const Duration(seconds: 2));

    _logInEmailIsLoading = false;
    setState(() {});

    // if (!mounted) return;
    // Navigator.of(context).pushNamedAndRemoveUntil(HomeScreen.routeName, (_) => false);
  }

  Future<void> _onRefresh() async {}
}
