import 'package:flutter/material.dart';
import 'package:rxdart/subjects.dart';

enum LanguageEnum { en, vi }

/// Language common bloc.
///
/// Usage: Call from Widgets
/// ```
/// final languageBloc = Provider.of<LanguageBloc>(context);
/// languageBloc.changeLanguage(LanguageEnum.vi);
/// languageBloc.changeLanguage(LanguageEnum.en);
/// ```
class LanguageBloc {
  final List<Locale> _supportedLocales = const [
    Locale('vi'),
    Locale('en'),
  ];
  Locale _locale = Locale(LanguageEnum.vi.name);

  /// Broadcast new language if it changes
  PublishSubject<Locale> languageChangedSubject = PublishSubject();

  /// Return a list of supported languages
  List<Locale> get supportedLocales => _supportedLocales;

  /// Return the current language, default to `Locale('vi')`;
  Locale get locale => _locale;

  void changeLanguage(LanguageEnum language) {
    if (language == LanguageEnum.en) {
      _locale = Locale(LanguageEnum.en.name);
      languageChangedSubject.add(_locale);
    } else {
      _locale = Locale(LanguageEnum.vi.name);
      languageChangedSubject.add(_locale);
    }
  }
}
