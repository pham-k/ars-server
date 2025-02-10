import 'package:ars_app/base/design/design.dart';
import 'package:ars_app/base/widget/button/button_filled.dart';
import 'package:ars_app/base/widget/button/button_outlined.dart';
import 'package:ars_app/base/widget/input/input.dart';
import 'package:flutter/material.dart';
import 'package:phone_form_field/phone_form_field.dart';
import 'package:provider/provider.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class Dev extends StatefulWidget {
  const Dev({super.key});

  @override
  State<Dev> createState() => _DevState();
}

class _DevState extends State<Dev> {
  bool _loading = false;
  late Design _des;
  late AppLocalizations _loc;

  @override
  Widget build(BuildContext context) {
    _des = Provider.of<Design>(context);
    _loc = AppLocalizations.of(context)!;
    return ListView(
      children: [
        Card.filled(
          child: Padding(
            padding: const EdgeInsets.all(16),
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: <Widget>[
                const ListTile(
                  leading: Icon(Icons.album),
                  title: Text('The Enchanted Nightingale'),
                  subtitle: Text('Music by Julie Gable. Lyrics by Sidney Stein.'),
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.end,
                  children: <Widget>[
                    ButtonOutlined(
                      label: "Test 1",
                      fullWidth: true,
                      onPressed: () {/* ... */},
                    ),
                    const SizedBox(width: 8),
                    ButtonFilled(
                      label: "Test 2",
                      fullWidth: true,
                      onPressed: () {/* ... */},
                    ),
                  ],
                ),
              ],
            ),
          ),
        ),
        // TextFormField(
        //   controller: TextEditingController(text: "TextFormField"),
        // ),
        // InputEmail(
        //   controller: TextEditingController(text: "InputEmail"),
        // ),
        // InputEmail(
        //   controller: TextEditingController(text: "InputEmail disabled"),
        //   enabled: false,
        // ),
        // InputEmail(
        //   controller: TextEditingController(text: "InputEmail readOnly"),
        //   readOnly: true,
        // ),
        // InputPassword(
        //   controller: TextEditingController(text: "InputPassword"),
        // ),
        // InputPassword(
        //   controller: TextEditingController(text: "InputPassword disabled"),
        //   enabled: false,
        // ),
        // InputPassword(
        //   controller: TextEditingController(text: "InputPassword readOnly"),
        //   readOnly: true,
        // ),
        // InputPhone(
        //   controller: PhoneController(initialValue: const PhoneNumber(isoCode: IsoCode.VN, nsn: "0909090909")),
        // ),
        // InputPhone(
        //   controller: PhoneController(initialValue: const PhoneNumber(isoCode: IsoCode.VN, nsn: "0909090909")),
        //   enabled: false,
        // ),
        // Text("Display Large", style: Theme.of(context).textTheme.displayLarge,),
        // Text("Display Medium", style: Theme.of(context).textTheme.displayMedium,),
        // Text("Display Small", style: Theme.of(context).textTheme.displaySmall,),
        // Text("Headline Large", style: Theme.of(context).textTheme.headlineLarge,),
        // Text("Headline Medium", style: Theme.of(context).textTheme.headlineMedium,),
        // Text("Headline Small", style: Theme.of(context).textTheme.headlineSmall,),
        // Text("Title Large", style: Theme.of(context).textTheme.titleLarge,),
        // Text("Title Medium", style: Theme.of(context).textTheme.titleMedium,),
        // Text("Title Small", style: Theme.of(context).textTheme.titleSmall,),
        // Text("Body Large", style: Theme.of(context).textTheme.bodyLarge,),
        // Text("Body Medium", style: Theme.of(context).textTheme.bodyMedium,),
        // Text("Body Small", style: Theme.of(context).textTheme.bodySmall,),
        // Text("Label Large", style: Theme.of(context).textTheme.labelLarge,),
        // Text("Label Medium", style: Theme.of(context).textTheme.labelMedium,),
        // Text("Label Small", style: Theme.of(context).textTheme.labelSmall,),
        // TextButton(onPressed: () {}, child: const Text("Text Button")),
        // ElevatedButton(onPressed: () {}, child: const Text("Elevated Button")),
        // OutlinedButton(onPressed: () {}, child: const Text("Outlined Button")),
        // IconButton(onPressed: () {}, icon: const Icon(Icons.add)),
        // ButtonFilled(
        //   label: "Button filled",
        //   onPressed: _onPressedButton,
        //   loading: _loading,
        // ),
        // ButtonFilled(
        //   label: "Button filled",
        //   onPressed: null,
        //   loading: _loading,
        // ),
        // ButtonFilled.icon(
        //   label: "Button filled icon",
        //   icon: const Icon(Icons.add),
        //   onPressed: _onPressedButton,
        //   loading: _loading,
        // ),
        // ButtonFilled.icon(
        //   label: "Button filled icon",
        //   icon: const Icon(Icons.add),
        //   onPressed: null,
        //   loading: _loading,
        // ),
        // ButtonOutlined(
        //   label: "Button outlined",
        //   onPressed: _onPressedButton,
        //   loading: _loading,
        // ),
        // ButtonOutlined(
        //   label: "Button outlined",
        //   onPressed: null,
        //   loading: _loading,
        // ),
        // ButtonOutlined.icon(
        //   label: "Button outlined icon",
        //   icon: const Icon(Icons.add),
        //   onPressed: _onPressedButton,
        //   loading: _loading,
        // ),
        // ButtonOutlined.icon(
        //   label: "Button outlined icon",
        //   icon: const Icon(Icons.add),
        //   onPressed: null,
        //   loading: _loading,
        // ),
      ],
    );
  }

  _onPressedButton() {
    setState(() {
      _loading = true;
    });

    Future.delayed(const Duration(seconds: 5)).then((_) {
      setState(() {
        _loading = false;
      });
    });
  }
}
