package main

import (
	"context"
	"fmt"
	"os"

	"github.com/brumhard/pr-dashboard/pkg/pr"
	"go.uber.org/zap"

	// This controls the maxprocs environment variable in container runtimes.
	// see https://martin.baillie.id/wrote/gotchas-in-the-go-network-packages-defaults/#bonus-gomaxprocs-containers-and-the-cfs
	_ "go.uber.org/automaxprocs"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	defer logger.Sync()

	github, err := pr.NewGitHubProvider(&pr.GitHubConfig{})
	if err != nil {
		return err
	}

	repos, err := github.ListReposForProject(context.Background(), "stackitcloud")
	if err != nil {
		return err
	}
	fmt.Println(repos)
	fmt.Println(len(repos))

	return nil
}
