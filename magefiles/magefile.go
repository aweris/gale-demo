//go:build mage

package main

import (
	"context"
	"os"

	"dagger.io/dagger"

	"github.com/magefile/mage/mg"

	"github.com/aweris/gale/pkg/gale"
	"github.com/aweris/gale/pkg/model"
)

type Go mg.Namespace

// All imports `ci` workflow's `build` job and add `test` and `lint` steps to it. Then runs the workflow.
func (_ Go) All(ctx context.Context) error {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}

	_, err = gale.New(client).
		WithJob("ci", "build").
		WithStep(&model.Step{ID: "0", Run: "echo 'Override checkout step'"}, true). // Override checkout step to execute current state of repository.
		WithStep(&model.Step{Run: "go test -v ./..."}, false).
		WithStep(&model.Step{Uses: "golangci/golangci-lint-action@v3"}, false).
		Exec(ctx)

	return err
}

// Lint runs linter on the project using gale without any workflow file.
func Lint(ctx context.Context) error {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}

	_, err = gale.New(client).
		WithStep(&model.Step{Uses: "actions/setup-go@v4", With: map[string]string{"go-version": "1.20"}}, false).
		WithStep(&model.Step{Uses: "golangci/golangci-lint-action@v3"}, false).
		Exec(ctx)

	return err
}

// Build builds the project using gale without any workflow file.
func Build(ctx context.Context) error {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}

	_, err = gale.New(client).
		WithStep(&model.Step{Uses: "actions/setup-go@v4", With: map[string]string{"go-version": "1.20"}}, false).
		WithStep(&model.Step{Run: "go build ./..."}, false).
		Exec(ctx)

	return err
}

// Test runs tests on the project using gale without any workflow file.
func Test(ctx context.Context) error {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}

	_, err = gale.New(client).
		WithStep(&model.Step{Uses: "actions/setup-go@v4", With: map[string]string{"go-version": "1.20"}}, false).
		WithStep(&model.Step{Run: "go test -v ./..."}, false).
		Exec(ctx)

	return err
}
