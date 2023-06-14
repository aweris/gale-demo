# gale-demo

This is a demo repository for [gale](https://github.com/aweris/gale) project. It contains hello world application
and a workflow that builds, tests and lint the application.

Demo has three different usage scenarios:

- Running workflow locally with gale cli
- Running workflow locally with mage and gale as a library
- Running workflow on github actions 

## Running workflow locally with gale cli

### Make targets

- `make build` - Builds the application
- `make test` - Runs the tests
- `make lint` - Runs the linter

### Usage:

```bash
make lint
```

This will download dagger and gale cli to `.bin` directory and run

```bash
./bin/dagger run ./bin/gale run ci lint --disable-checkout --export
```

This will execute the `lint` job in the workflow `.github/workflows/ci.yaml`. In executed command,

- `ci` is the name of the workflow. If workflow doesn't have name, than it must be relative path to the workflow file
- `lint` is the name of the job
- `--disable-checkout` flag is a temporary testing flag overrides first step of the workflow which is always checkout to test current code. (**Note**: This flag will be removed in the future, it's just a hack for now)
- `--export` flag exports the runner directory after the execution. Exported directory will be placed under `.gale` directory in the current directory.

## Running workflow locally with mage and gale as a library

### Make targets

This scenario uses mage as a build tool and gale as a library. targets are same as above but they are prefixed with `mage.` 
and only difference is they use `gale` as a library instead of `gale` cli.

- `make mage.build` - Runs `mage build` command to build the application.
- `make mage.test` - Runs `mage test` command to run the tests.
- `make mage.lint` - Runs `mage lint` command to run the linter.
- `make mage.all` - Runs `mage go:all` command to run all targets above. This is basically uses `ci/build` job in 
and adds `test` and `lint` steps to customize the run.

### Usage:

```go
// All imports `ci` workflow's `build` job and add `test` and `lint` steps to it. Then runs the workflow.
func (_ Go) All(ctx context.Context) error {
    client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
    if err != nil {
    return err
    }
    
    _, err = gale.New(client). // Create a new gale instance with default container and given dagger client.
        WithJob("ci", "build"). // Add `build` job from `ci` workflow to configuration.
        WithStep(&model.Step{ID: "0", Run: "echo 'Override checkout step'"}, true). // Override checkout step to execute current state of repository. This is assumes checkout step is the first step in the job without any id.
        WithStep(&model.Step{Run: "go test -v ./..."}, false). // Adds a new step to run tests.
        WithStep(&model.Step{Uses: "golangci/golangci-lint-action@v3"}, false). // Adds a new step to run linter.
        Exec(ctx) // Executes the configuration.
    
    return err
}
```

With this code, 

- Importing a job from a workflow
- Overriding a step in the job
- Adding new steps to the job
- Running the workflow


## Running workflow on github actions

This scenario uses github actions to run the gale. The workflow is defined in `.github/workflows/gale-on-gha.yml` file. This
usage allow us loading github action configuration from runner to execute our configuration in the dagger container.

Workflow:

```yaml
name: gale on gha

on: workflow_dispatch

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v3

      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.20'

      - run: make mage.lint
        env:
          GH_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

mage lint:

```go
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
```

This workflow/mage snippet is,

- Checkout the repository
- Setup go
- Run `make mage.lint` command which is wrapper for `mage lint` command mentioned above.