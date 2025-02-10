import 'package:ars_app/base/design/design.dart';
import 'package:ars_app/base/widget/ars_logo/ars_logo.dart';
import 'package:ars_app/base/widget/top_app_bar/top_app_bar_button.dart';
import 'package:ars_app/base/widget/top_app_bar/top_app_bar_widget.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class MenuScreen extends StatefulWidget {
  const MenuScreen({super.key});

  static const routeName = '/menu';

  @override
  State<MenuScreen> createState() => _MenuScreenState();
}

class _MenuScreenState extends State<MenuScreen> {
  late Design _des;
  late AppLocalizations _loc;

  final List quotes = [
    {
      "quote":
      "It’s your place in the world; it’s your life. Go on and do all you can with it, and make it the life you want to live.",
      "author": "Mae Jemison"
    },
    {
      "quote":
      "You may be disappointed if you fail, but you are doomed if you don’t try.",
      "author": "Beverly Sills"
    },
    {
      "quote":
      "Remember no one can make you feel inferior without your consent.",
      "author": "Eleanor Roosevelt"
    },
    {
      "quote": "Life is what we make it, always has been, always will be.",
      "author": "Grandma Moses"
    },
    {
      "quote":
      "The question isn’t who is going to let me; it’s who is going to stop me.",
      "author": "Ayn Rand"
    },
    {
      "quote":
      "When everything seems to be going against you, remember that the airplane takes off against the wind, not with it.",
      "author": "Henry Ford"
    },
    {
      "quote":
      "It’s not the years in your life that count. It’s the life in your years.",
      "author": "Abraham Lincoln"
    },
    {
      "quote": "Change your thoughts and you change your world.",
      "author": "Norman Vincent Peale"
    },

  ];

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
    );
  }

  PreferredSize _buildAppBar() {
    return PreferredSize(
      preferredSize: const Size(double.infinity, 40),
      child: TopAppBar(
        title: _loc.menu_screen_title,
        leading: TopAppBarButton(icon: TopAppBarIcon.chevronLeft, onTap: _onTapBack),
      ),
    );
  }

  Widget _buildBody() {
    return Column(
      children: [
        Expanded(
          child: Card.filled(
            child: ListView.builder(
              itemCount: quotes.length,
              itemBuilder: (BuildContext context, int index) {
                return _buildListItemExpandable(quotes[index]);
              },
            ),
          ),
        ),
        Expanded(
          child: Card.filled(
            child: ListView.builder(
              itemCount: quotes.length,
              itemBuilder: (BuildContext context, int index) {
                return _buildListItemExpandable(quotes[index]);
              },
            ),
          ),
        )
      ],
    );
  }

  Widget _buildListItem(item) {
    return ListTile(
      title: Text(
        item['quote'],
        style: _des.typo.bodyMedium,
      ),
    );
  }

  Widget _buildListItemExpandable(item) {
    return ExpansionTile(
      title: Text(
        item['author'],
      ),
      children: <Widget>[
        ListTile(
          title: Text(
            item['quote'],
            style: _des.typo.bodyMedium,
          ),
        )
      ],
    );
  }

  void _onPopInvokedWithResult (bool didPop, result) {}

  void _onTapBack() {
    Navigator.of(context).pop();
  }
}
