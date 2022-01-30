import 'package:app/services/pr.dart';
import 'package:app/services/theme.dart';
import 'package:app/ui/animate_once.dart';
import 'package:app/ui/colors.dart';
import 'package:app/ui/repo_card.dart';
import 'package:app/ui/skeletons/skeleton_repo_card.dart';
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:provider/provider.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        ChangeNotifierProvider<ThemeService>(create: (_) => ThemeService()),
        Provider<PrService>(
          create: (_) => PrService(endpoint: "localhost"),
          dispose: (_, prService) => prService.dispose(),
        ),
      ],
      child: Consumer<ThemeService>(
        builder: (context, themeService, child) {
          return MaterialApp(
            title: 'pr:mate',
            theme: themeService.theme,
            home: const Home(title: 'pr:mate'),
          );
        },
      ),
    );
  }
}

class Home extends StatelessWidget {
  const Home({Key? key, required this.title}) : super(key: key);
  final String title;

  @override
  Widget build(BuildContext context) {
    PrService prService = Provider.of<PrService>(context, listen: false);

    return Scaffold(
      appBar: AppBar(
        title: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            IconButton(
              onPressed: () => prService.loadPRs(),
              icon: TurnAtLeastOnceOnListenable(
                child: const Icon(Icons.replay_rounded),
                listenable: prService.isLoading,
              ),
            ),
            Text(
              title,
              style: GoogleFonts.novaMono(
                  textStyle: Theme.of(context).textTheme.headline5),
            ),
            Container()
          ],
        ),
        flexibleSpace: Container(
          decoration: const BoxDecoration(
            gradient: LinearGradient(
              begin: Alignment.topLeft,
              end: Alignment.bottomRight,
              colors: [pink, orange],
            ),
          ),
        ),
      ),
      body: Center(
        child: ConstrainedBox(
          constraints: const BoxConstraints(maxWidth: 1200),
          child: Center(
            child: StreamBuilder<List<Repository>?>(
              initialData: null,
              stream: prService.stream,
              builder: (context, snapshot) {
                if (snapshot.hasError) {
                  // fix for showing snackbar during builder:
                  // https://stackoverflow.com/questions/54230331/what-is-the-fancy-way-to-use-snackbar-in-streambuilder
                  WidgetsBinding.instance!.addPostFrameCallback((_) {
                    ScaffoldMessenger.of(context).showSnackBar(
                      SnackBar(
                        content: Text(snapshot.error.toString()),
                        backgroundColor: Theme.of(context).colorScheme.primary,
                      ),
                    );
                  });
                }

                if (snapshot.connectionState == ConnectionState.waiting ||
                    !snapshot.hasData) {
                  return ListView(
                    children: const [
                      RepositoryCardSkeleton(),
                      RepositoryCardSkeleton()
                    ],
                  );
                }

                List<Repository> repos = snapshot.data!;
                repos.sort((a, b) =>
                    b.pullrequests.length.compareTo(a.pullrequests.length));
                return ListView(
                  children: repos
                      .map((repo) => RepositoryCard(
                            repo: repo,
                          ))
                      .toList(),
                );
              },
            ),
          ),
        ),
      ),
      floatingActionButton: Consumer<ThemeService>(
        builder: (context, themeService, child) {
          return FloatingActionButton(
            backgroundColor: Theme.of(context).colorScheme.primary,
            onPressed: () => {
              themeService.setTheme(
                  themeService.currentBrightness == Brightness.light
                      ? Brightness.dark
                      : Brightness.light)
            },
            child: Icon(themeService.currentBrightness == Brightness.light
                ? Icons.nightlight_round
                : Icons.wb_sunny),
          );
        },
      ),
    );
  }
}
