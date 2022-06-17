import 'package:flutter/material.dart';

class Bye extends StatelessWidget {
  const Bye({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Column(
            children: [
              Text(
                "Heute gibts keine PRs zu reviewen.",
                style: Theme.of(context).textTheme.headline3,
              ),
              Text(
                "Heute wird einfach lässig reingeloggert.",
                style: Theme.of(context).textTheme.headline3,
              ),
              Text(
                "🚀 Haut rein 🚀",
                style: Theme.of(context).textTheme.headline3,
              ),
            ],
          ),
          GestureDetector(
            child: Image.network(
              'https://c.tenor.com/Pm4S40MGsIQAAAAC/hacker-hackerman.gif',
              width: 600,
            ),
            onTap: () {
              WidgetsBinding.instance!.addPostFrameCallback((_) {
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(
                    content: Center(
                        child: Text(
                      "💥BOOM💥",
                      style: Theme.of(context).textTheme.headline1,
                    )),
                    backgroundColor: Theme.of(context).colorScheme.primary,
                  ),
                );
              });
            },
          ),
        ],
      ),
    );
  }
}
