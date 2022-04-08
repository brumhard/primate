# `pr:mate`

[![GoTemplate](https://img.shields.io/badge/go/template-black?logo=go)](https://github.com/SchwarzIT/go-template)

Pluggable dashboard to make PRs from multiple sources on multiple repos visible.

> This repo is currently WIP. Further docs will follow.

## Development

The project uses `make` to make your life easier. If you're not familiar with
Makefiles you can take a look at [this quickstart guide](https://makefiletutorial.com).

Whenever you need help regarding the available actions, just use the following
command.

```shell
make help
```

### Setup

To get your setup up and running do the following.

1. Install [dart and flutter](https://docs.flutter.dev/get-started/install)
2. Install `protoc-gen-dart`:

   ```shell
   flutter pub global activate protoc_plugin
   ```
   
   This command recommends to add a location to your `PATH`, do this as
   instructed. (See also the next step.)

   It might be necessary to also install `protoc_plugin` via `dart` directly, by
   running:

   ```shell
   dart pub global activate protoc_plugin
   ```
  
3. Add  the `.pub-cache/bin` location to your `PATH`:

   ```shell
   export PATH="$PATH:$HOME/.pub-cache/bin"
   # or if on flutter it could also be sth like
   export PATH="$PATH":"$HOME/sdk/flutter/.pub-cache/bin"
   ```

4. Setup everything else:

   ```shell
   make all
   ```

   This will initialize a git repo, download the dependencies in the latest
   versions and install all needed tools. If needed code generation will be
   triggered in this target as well.

5. Add configuration for the backend service:

   ```shell
   cp configs/config_example.yaml config/config.yaml
   vi config/config.yaml
   ```

   `pr:mate` supports GitHub, Azure DevOps and BitBucket repositories. You can
   also use wildcards for repositories, like so:

   ```yaml
   providers:
   - providerType: "github"
     repositories:
       - "brumhard/*"
     extraConfig:
       pat: "xxxxxxxxxx"
   ```

   The `pat` (Personal Access Token) needs to be created with the correct
   permissions.
   
   * **GitHub**: See [the example for GitHub](./docs/img/github_token.png).
   * **Azure DevOps**: Your PAT needs the `Code (Read)` permission.
   * **BitBucket**: TODO

You're now ready to start developing!

### Running

When developing the frontend, use regular `flutter` commands in the `app`
folder:

```shell
cd app && flutter run
```

In order to have a backend service to develop against, use the following:

```shell
# build the backend
make build

# start the server
./out/bin/primate --config=configs/config.yaml

# in a separate terminal, start developing the frontend
flutter run
```

### Test & lint

Run linting

```bash
make lint
```

Run tests

```bash
make test
```
