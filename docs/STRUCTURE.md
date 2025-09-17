# Project Structure

```text
./
|-- cmd/
|   |-- awscbtui/
|       |-- main.go
|
|-- pkg/
|   |-- ui/
|   |   |-- app.go
|   |   |-- layout.go
|   |   |-- keymap.go
|   |   |-- navigation.go
|   |   |-- render/
|   |       |-- projects.go
|   |       |-- builds.go
|   |       |-- details.go
|   |       |-- logs.go
|   |       |-- status.go
|   |       |-- help.go
|   |
|   |-- state/
|   |   |-- state.go
|   |   |-- intents.go
|   |   |-- reducer.go
|   |   |-- selectors.go
|   |   |-- messages.go
|   |
|   |-- jobs/
|   |   |-- dispatcher.go
|   |   |-- types.go
|   |   |-- timers.go
|   |
|   |-- aws/
|   |   |-- session.go
|   |   |-- codebuild.go
|   |   |-- logs.go
|   |   |-- types.go
|   |   |-- fake/
|   |       |-- codebuild_fake.go
|   |       |-- logs_fake.go
|   |
|   |-- config/
|   |   |-- load.go
|   |   |-- appconfig.go
|   |
|   |-- logging/
|   |   |-- commandlog.go
|   |   |-- slog_adapter.go
|   |
|   |-- util/
|   |   |-- ringbuffer.go
|   |   |-- debounce.go
|   |   |-- backoff.go
|   |   |-- copy.go
|
|-- test/
|   |-- reducer_test.go
|   |-- intents_test.go
|   |-- aws_fakes_test.go
|   |-- ui_headless_test.go
|
|-- docs/
|   |-- DESIGN.md
|   |-- PLAN.md
|   |-- STRUCTURE.md
|
|-- go.mod
|-- go.sum
|-- README.md
```


Here are AI notes on the structure and what each part is for, plus some Go tips as you work.

Top level

    go.mod / go.sum: Tracks your module name and dependencies. Always run go mod tidy after adding imports to keep these clean.
    README.md: What the project does and how to run it. Keep it short but useful.
    docs/: Longer docs like design, plan, and structure. Useful for future you.

cmd/awscbtui/

    main.go: The entry point. Keep it tiny—just create the app and run it.
    Rule of thumb: cmd/<appname> contains small wrappers; all logic lives under pkg/.

pkg/ui/

    app.go: Creates the gocui Gui, holds the main loop, wires everything together. One place to initialize state, dispatcher, and keybindings.
    layout.go: Defines panels/views (projects, builds, bottom tabs, status bar). Only draw; do not fetch data here.
    keymap.go: Maps keys (e.g., j/k/r/?) to “intents” (events). Keep it declarative and consistent.
    navigation.go: Tracks which panel is focused and how Tab/Shift-Tab move focus.
    render/: Pure functions that take state and write strings to views. They should not call AWS or spawn goroutines. Keep rendering pure and predictable.

Why: Separating UI rendering from logic makes it easier to test and reason about. If the UI isn’t updating, you check state first, then rendering.

pkg/state/

    state.go: The AppState struct—your single source of truth. Contains session, UI focus, caches (projects, builds), etc.
    intents.go: Definitions of user actions and async events (e.g., RefreshProjects, BuildsLoaded). Usually a small enum + payload structs.
    reducer.go: Pure functions that take (state, intent) and return a new or mutated state. This keeps logic predictable and testable.
    selectors.go: Helpers to read derived values from state (e.g., current selected project). Keeps UI code simple.
    messages.go: Defines the messages your jobs send back (e.g., Loaded, Error, PartialResult).

Why: Treat state transitions as “reducers” (inspired by Redux). Easier to test: given input state + intent => expected state.

pkg/jobs/

    dispatcher.go: Runs background jobs (goroutines) and sends results back via channels. Handles deduping, debouncing, and cancellation.
    types.go: Keys and metadata to identify jobs (e.g., {Service: CodeBuild, Scope: ProjectName}).
    timers.go: Tickers for auto-refresh logic.
    Tip: Always pass context.Context to jobs so you can cancel them when profile/region/focus changes.

Why: Never block the UI. Offload network calls to workers, and communicate via channels.

pkg/aws/

    session.go: Loads AWS config/credentials (respects env vars, profiles), exposes an AWS config object and identity info. Handle SSO prompts here.
    codebuild.go: Thin wrappers around AWS SDK v2 calls for CodeBuild: ListProjects, ListBuildsForProject, BatchGetBuilds. Convert SDK types to your own DTOs.
    logs.go: CloudWatch Logs tailer logic (GetLogEvents with nextToken). Expose a function that returns new lines and a next cursor.
    types.go: Your DTOs (Project, Build, BuildPhase, LogLine). Keep these simple Go structs.
    fake/: Fake implementations used by tests so you don’t hit AWS during unit tests.

Why: Isolate the SDK. If the SDK changes, only this folder changes. The rest of the app uses your stable DTOs.

pkg/config/

    load.go: Loads ~/.aws for auth (via the SDK) and your app’s own config file (e.g., ~/.config/awscbtui/config.yml).
    appconfig.go: Struct that holds app settings (keybindings override, refresh intervals).
    Tip: Provide sensible defaults if no config file exists.

pkg/logging/

    commandlog.go: A thread-safe buffer storing recent “commands” your app executed (e.g., CodeBuild.ListProjects … took 120ms). The UI reads this for the command log panel.
    slog_adapter.go: Optional integration with Go’s slog or other loggers for debugging.

pkg/util/

    ringbuffer.go: A fixed-size buffer for log tail lines so memory doesn’t grow forever.
    debounce.go: Utility to delay jobs slightly to avoid spamming AWS on rapid keypresses.
    backoff.go: Exponential backoff with jitter to handle throttling/network errors gracefully.
    copy.go: Cross-platform clipboard helper (optional nice-to-have).

test/

    reducer_test.go: Tests for state transitions—fast and reliable.
    intents_test.go: Tests that intents do the right thing to the state.
    aws_fakes_test.go: Tests for your AWS wrappers using fakes (no real AWS calls).
    ui_headless_test.go: Optional; simulate key presses and assert state changes without drawing the actual UI.
