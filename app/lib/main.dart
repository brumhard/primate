import 'package:app/services/pr.dart';
import 'package:app/ui/repo_card.dart';
import 'package:app/ui/skeletons/skeleton_repo_card.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
        title: 'PullRequest Dashboard',
        theme: ThemeData(
          primarySwatch: Colors.blue,
        ),
        home: MultiProvider(
          providers: [
            Provider<PrService>(
              create: (context) => PrService(endpoint: "localhost"),
            )
          ],
          child: const Home(title: 'PullRequest Dashboard'),
        ));
  }
}

class Home extends StatelessWidget {
  const Home({Key? key, required this.title}) : super(key: key);
  final String title;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(title),
        flexibleSpace: Container(
          decoration: const BoxDecoration(
            gradient: LinearGradient(
              begin: Alignment.topLeft,
              end: Alignment.bottomRight,
              colors: [Color(0xFFdd5e89), Color(0xFFf7bb97)],
            ),
          ),
        ),
      ),
      body: Center(
        child: ConstrainedBox(
          constraints: const BoxConstraints(maxWidth: 1200),
          child: FutureBuilder(
            future: Provider.of<PrService>(context, listen: false).getAllPRs(),
            builder: (context, AsyncSnapshot<List<Repository>> snapshot) {
              return Center(
                child: widgetsForSnapshotState(snapshot),
              );
            },
          ),
        ),
      ),
    );
  }
}

Widget widgetsForSnapshotState(AsyncSnapshot<List<Repository>> snapshot) {
  if (snapshot.hasError) {
    // TODO: show snackbar instead
    return Column(children: [
      const Icon(
        Icons.error_outline,
        color: Colors.red,
        size: 60,
      ),
      Padding(
        padding: const EdgeInsets.only(top: 16),
        child: Text('Error: ${snapshot.error}'),
      ),
    ]);
  }
  if (!snapshot.hasData) {
    return ListView(
      children: const [RepositoryCardSkeleton(), RepositoryCardSkeleton()],
    );
  }

  snapshot.data!
      .sort((a, b) => b.pullrequests.length.compareTo(a.pullrequests.length));
  return ListView(
    children: snapshot.data!
        .map((repo) => RepositoryCard(
              repo: repo,
            ))
        .toList(),
  );
}
