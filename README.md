# awscbtui

awscbtui is a terminal UI for browsing AWS CodeBuild projects, builds, and live CloudWatch Logs. It provides a fast, keyboard-driven interface to inspect builds, tail logs, and run common actions without leaving the terminal.

## Features

- Browse CodeBuild projects and builds
- View build details (phases, environment, artifacts)
- Tail CloudWatch Logs for a selected build with efficient buffering
- Command log showing AWS calls and timings
- Configurable keybindings and refresh intervals
- Fake AWS implementations for fast unit tests

## Project structure

- `cmd/awscbtui`: CLI entrypoint (main.go)
- `pkg/ui`: TUI code (layout, renderers, keymap, navigation)
- `pkg/state`: AppState, intents, reducer, selectors
- `pkg/jobs`: Background job dispatcher, timers, job types
- `pkg/aws`: AWS SDK wrappers and DTOs; `fake/` contains test fakes
- `pkg/config`: Config loading and app settings
- `pkg/logging`: Command log and logger adapter
- `pkg/util`: Utilities (ring buffer, debounce, backoff, clipboard)
- `test/`: Unit tests for reducers, intents, and fakes
- `docs/`: Design, PLAN, STRUCTURE, and NOTES

## Quick start

Prerequisites
- Go 1.20+
- AWS credentials/config in `~/.aws` or environment variables for real AWS access

Build

```bash
go build -o bin/awscbtui ./cmd/awscbtui
```

Run

```bash
# run with real AWS credentials
./bin/awscbtui

# run with fakes (TBD: feature flag to use fakes)
./bin/awscbtui --use-fakes
```

## Configuration

The app loads AWS config from the environment and `~/.aws`. App-specific config can be placed in `~/.config/awscbtui/config.yml` (example config coming soon). Precedence: environment > AWS config > app config.

## Development

- Keep UI updates on the main goroutine; use channels for worker results.
- Use contexts for cancellable background jobs.
- Write reducer tests first for predictable state management.
- Use fakes in `pkg/aws/fake` for fast unit tests without hitting AWS.

Run tests

```bash
go test ./...
```

Linting & formatting

```bash
gofmt -w .
# Optional: golangci-lint run
```

## Roadmap

See `docs/PLAN.md` for milestone breakdown (skeleton, AWS integration, logs tailing, UX polish).
