import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:ars_app/base/design/design.dart';
import 'package:ars_app/screen/splash/splash_screen.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:provider/provider.dart';
import 'package:provider/single_child_widget.dart';
import 'package:rxdart/rxdart.dart';

import 'base/bloc/language_bloc.dart';
import 'base/design/material_theme/material_theme.dart';
import 'routes.dart';

Future<void> main() async {
  runApp(const ArsApp());
}

class ArsApp extends StatefulWidget {
  const ArsApp({super.key});

  @override
  State<StatefulWidget> createState() => _ArsAppState();
}

class _ArsAppState extends State<ArsApp> {
  late List<SingleChildWidget> _providers;

  final String _initialRoute = SplashScreen.routeName;
  late Routes _routes;

  late LanguageBloc _languageBloc;
  late Locale _locale;

  late Design _ds;
  ArsTheme _theme = ArsTheme.light;
  final PublishSubject<ArsTheme> _onThemeChanged = PublishSubject<ArsTheme>();

  @override
  void initState() {
    super.initState();
    initPlatformState();

    SystemChrome.setPreferredOrientations([
      DeviceOrientation.portraitUp,
    ]);

    _routes = Routes();

    _languageBloc = LanguageBloc();
    _locale = _languageBloc.locale;
    _languageBloc.languageChangedSubject.listen((Locale locale) {
      setState(() {
        _locale = locale;
      });
    });

    _ds = Design(
      theme: _theme,
      onThemeChanged: _onThemeChanged
    );


    _providers = [
      Provider<LanguageBloc>.value(value: _languageBloc),
      Provider<Design>.value(value: _ds),
    ];
  }

  @override
  void dispose() {
    super.dispose();
  }

  Future<void> initPlatformState() async {}

  @override
  Widget build(BuildContext context) {
    return MultiProvider(
        providers: _providers,
        child: MaterialApp(
          navigatorObservers: const [],
          locale: _locale,
          localizationsDelegates: const [
            AppLocalizations.delegate,
            GlobalMaterialLocalizations.delegate,
            GlobalWidgetsLocalizations.delegate,
            GlobalCupertinoLocalizations.delegate,
          ],
          supportedLocales: _languageBloc.supportedLocales,
          initialRoute: _initialRoute,
          routes: _routes.routes,
          theme: MaterialTheme(_ds).themeData,
          darkTheme: MaterialTheme(_ds).themeData,
          highContrastTheme: MaterialTheme(_ds).themeData,
          highContrastDarkTheme: MaterialTheme(_ds).themeData,
          themeMode: _ds.themeMode,
          builder: (context, child) {
            return MediaQuery(
              data: MediaQuery.of(context)
                  .copyWith(textScaler: const TextScaler.linear(1.0)),
              child: child!,
            );
          },
        ));
  }

  void _listenOnThemsChanged() {
    _onThemeChanged.listen((ArsTheme theme) {
      setState(() {
        _theme = theme;
      });
    });
  }
}
