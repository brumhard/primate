# `pr:mate`

[![GoTemplate](https://img.shields.io/badge/go/template-black?logo=go)](https://github.com/SchwarzIT/go-template)

Pluggable dashboard to make PRs from multiple sources on multiple repos visible.

The project uses `make` to make your life easier. If you're not familiar with Makefiles you can take a look at [this quickstart guide](https://makefiletutorial.com).

Whenever you need help regarding the available actions, just use the following command.

```bash
make help
```

## Setup

To get your setup up and running do the following.

- install dart and flutter
- install protoc-gen-dart
  ```shell
  flutter pub global activate protoc_plugin
  ```
- add pub cache bin folder to `PATH`
  ```shell
  export PATH="$PATH:$HOME/.pub-cache/bin"
  # or if on flutter it could also be sth like
  export PATH="$PATH":"$HOME/sdk/flutter/.pub-cache/bin"
  ```
- setup everything else:
```bash
make all
```

This will initialize a git repo, download the dependencies in the latest versions and install all needed tools.
If needed code generation will be triggered in this target as well.

## Test & lint

Run linting

```bash
make lint
```

Run tests

```bash
make test
```
