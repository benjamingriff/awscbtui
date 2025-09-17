Design Document: AWS CodeBuild TUI (Lazygit-Inspired)

Vision and Scope

- Purpose: A fast, keyboard-first terminal app to monitor AWS CodeBuild projects and builds without opening the AWS Console.

- Scope (initial): Read-only monitoring
	- List CodeBuild projects

	- List recent builds per project

	- Live-tail a selected build’s logs

	- Show build status, phase transitions, durations, and basic artifacts info


- Non-goals (for now): Start/stop builds, mutate resources, cross-service orchestration.

Product Principles (borrowed from Lazygit)

- Single event loop: All UI updates flow from a central state store.

- Single source of truth: State struct holds truth; views render from state.

- Reactive UI: Keypresses and async events create “intents,” which update state and trigger re-renders.

- Async-aware: Never block the UI; AWS calls run in goroutines with cancellation and backoff.

- Progressive disclosure: Core info is always visible; details and logs revealed on demand.

- Opinionated keyboard UX: Few, consistent keys; discoverable with “?” help panel.

- Debuggable by users: A visible “command log” shows AWS calls and outcomes.

Core User Journeys

- Browse projects
	- Open app -> see profile/region, projects in left panel, status bar with counts.


- Inspect builds for a project
	- Select project -> right panel shows recent builds with status, time, and build number.


- Tail a build’s logs
	- Select a build -> press “t” to tail

	- Bottom panel shows CloudWatch Logs streaming; phase changes highlighted.


- Refresh and filter
	- Press “r” to refresh focused list; “/” to filter by name or status.


UI Layout (gocui)

- Top status bar: Current AWS profile/region, identity, last refresh time, network status.

- Left panel (Projects): List of CodeBuild project names with success/failure counts.

- Right panel (Builds): For selected project, shows recent builds (id, status, duration, start time).

- Bottom panel (Logs/Details/Command Log):
	- Tabs: [Logs] [Details] [Cmd]

	- Logs: live tail for selected build

	- Details: JSON-ish formatted build description (phases, environment, artifacts)

	- Cmd: human-readable list of AWS calls executed


- Help overlay: “?” shows keybindings and shortcuts.

State Model (simplified)

- Session
	- profile (string), region (string), identity (ARN), ssoStatus, expiresAt


- Focus/UI
	- focusedPanel (projects|builds|logs|help), selection indices, activeTab

	- filterText, showHelp, notifications (toasts)


- Data caches
	- projects []Project

	- buildsByProject map[string][]Build

	- logTail map[buildID]RingBuffer

	- lastUpdated per collection


- Async/Jobs
	- inflight map[string]JobMeta (deduping and debouncing)

	- backoff per job key


- Keymaps
	- Defaults + user overrides


Types are normalized to thin DTOs detached from SDK structs. Convert SDK outputs once, then work with your DTOs in state and UI.

Intents (examples)

- Session/Config
	- ChangeProfile(p), ChangeRegion(r), StartSSOLogin


- Navigation
	- FocusNext, FocusPrev, OpenHelp, CloseHelp

	- ApplyFilter(text), ClearFilter


- Data
	- RefreshProjects, RefreshBuilds(projectName), RefreshBuild(buildID)

	- TailLogs(buildID), StopTail(buildID)


- UX
	- CopyBuildID, CopyProjectName


- Internals (async)
	- ProjectsLoaded, BuildsLoaded, BuildUpdated, LogsAppended, JobError


Each intent handler:


- Validates preconditions

- Mutates state minimally (pure-ish reducer style)

- Dispatches async work via dispatcher (goroutine + context)

- Emits command log entry

AWS Integration (Go SDK v2)

- Auth/session
	- Load config from ~/.aws/config and credentials (respect env overrides)

	- Surface SSO status; guide user if expired


- CodeBuild APIs
	- ListProjects (paginate)

	- BatchGetProjects (for details, optional)

	- ListBuildsForProject (paginate by time desc)

	- BatchGetBuilds (resolve statuses/phases)


- Logs
	- Build logs location: CloudWatch Logs group/stream from build’s logs info

	- CloudWatch Logs GetLogEvents or FilterLogEvents

	- Tail strategy: remember nextToken/timestamp to avoid duplicates


- Rate limiting
	- Use SDK retryer; add app-level jitter/backoff for refreshers

	- Avoid hammering: default refresh intervals (e.g., 15–30s on builds list, 1–2s for tail)


Refresh Strategy

- Manual refresh key “r” always available

- Auto-refresh:
	- Projects: 60s

	- Builds for selected project: 15s while panel focused, 60s otherwise

	- Tail: poll every 1–2s using nextToken; backoff on errors up to 30s


- Cancel on context change:
	- When profile/region changes, cancel all workers and clear caches


Error Handling and UX

- Distinguish auth vs permission vs throttling vs network

- Show actionable hints: “Press l to login via SSO”, “Try r to retry”

- Display error in command log and as brief toast

- For throttling, show “Backoff 10s” countdown in status bar

Keyboard Map (initial)

- Global
	- Tab / Shift-Tab: cycle panels

	- ?: help, q: quit

	- /: filter (Esc to clear)

	- r: refresh focused

	- p: profile menu; R: region menu

	- l: SSO login


- Projects panel
	- j/k: move; Enter: load builds for selected project


- Builds panel
	- j/k: move; Enter: open details tab

	- t: tail logs; s: stop tail

	- y: copy build id


- Logs/Details/CommandLog
	- h/l: switch tabs (Logs/Details/Cmd)


Testing Strategy

- Unit
	- Reducers/intent handlers (pure state transformations)

	- AWS wrappers using interface fakes (no real AWS)


- Integration
	- Headless UI loop: feed key events; assert state changes

	- Localstack optional for CodeBuild/Logs; or thin fakes with recorded responses


- Golden tests
	- Render snippets for rows formatting (optional, nice-to-have)


What to Defer

- Mutations (start/stop builds)

- Multi-account fans out

- Search across all builds by commit id (later)

- Complex filters and saved views

- Export to files
