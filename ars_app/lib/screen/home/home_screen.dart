
import 'package:ars_app/base/design/design.dart';
import 'package:ars_app/base/widget/ars_logo/ars_logo.dart';
import 'package:ars_app/base/widget/top_app_bar/top_app_bar.dart';
import 'package:ars_app/base/widget/top_app_bar/top_app_bar_button.dart';
import 'package:ars_app/screen/home/widget/feature/feature.dart';
import 'package:ars_app/screen/menu/menu_screen.dart';
import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:provider/provider.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  static const routeName = '/home';

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  late Design _des;
  late AppLocalizations _loc;

  @override
  Widget build(BuildContext context) {
    _des = Provider.of<Design>(context);
    _loc = AppLocalizations.of(context)!;

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
          appBar: _buildAppBar(),
          body: SafeArea(
            child: CustomScrollView(
              slivers: [
                SliverFillRemaining(
                  hasScrollBody: true,
                  child: Padding(
                    padding: EdgeInsets.all(_des.spacing.s(16)),
                    child: body,
                  ),
                )
              ],
            ),
          ),
        ),
      ),
    );
  }

  PreferredSize _buildAppBar() {
    return PreferredSize(
        preferredSize: const Size(double.infinity, 40),
        child: TopAppBar(
          title: _loc.home_screen_title,
          leading: Center(child: ArsLogo(size: _des.spacing.s(28),),),
          actions: [
            TopAppBarButton(icon: TopAppBarIcon.menu, onTap: _onTapMenu,)
          ],
        ),
    );
  }

  Widget _buildBody() {
    return Column(
      children: [
        _buildFeatureCashFlowCalculator(),
        _buildFeatureTodo(),
        _buildFeatureDocument(),
        _buildFeatureCalendar(),
      ],
    );
  }

  Widget _buildFeatureCashFlowCalculator() {
    return Feature(
      name: _loc.home_screen_cash_flow_calc,
      image: 'assets/images/svg/cash_flow.svg',
      onTap: _onTapCashFlowCalc,
    );
  }

  Widget _buildFeatureTodo() {
    return Feature(
      name: _loc.home_screen_to_do,
      image: 'assets/images/svg/todo.svg',
      onTap: _onTapToDo,
    );
  }

  Widget _buildFeatureDocument() {
    return Feature(
      name: _loc.home_screen_document,
      image: 'assets/images/svg/document.svg',
      onTap: _onTapDocument,
    );
  }

  Widget _buildFeatureCalendar() {
    return Feature(
      name: _loc.home_screen_calendar,
      image: 'assets/images/svg/calendar.svg',
      onTap: _onTapCalendar,
    );
  }

  void _onPopInvokedWithResult (bool didPop, result) {}

  Future<void> _onRefresh() async {}

  void _onTapMenu() {
    Navigator.of(context).pushNamed(MenuScreen.routeName);
  }

  void _onTapCashFlowCalc() {}

  void _onTapToDo() {}

  void _onTapDocument() {}

  void _onTapCalendar() {}

}
