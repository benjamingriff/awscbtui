# Implementation Plan (Milestones with Quick Wins)

## Week 1: Skeleton + State + Layout
- Implement gocui layout with 3 panels: Projects (left), Builds (right), Bottom (placeholder)
- Wire event loop, keybindings for navigation, help overlay
- Create AppState, intents, dispatcher stubs, and command log
- Quick win: render sample fake data in panels; keys move selection; help works

## Week 2: AWS Session + List Projects
- Add AWS SDK v2, config loading from ~/.aws and env
- Show current profile/region in status bar; simple profile switching (env override or menu)
- Implement ListProjects with pagination; normalize to []Project
- RefreshProjects intent and async job
- Quick win: See real projects listed in the left panel; command log shows AWS calls

## Week 3: List Builds for Project
- Implement ListBuildsForProject + BatchGetBuilds to enrich data
- Builds panel populates on project selection
- Add filtering “/” for project or build id/status
- Quick win: Real builds appear; “r” refreshes focused builds list

## Week 4: Build Details View
- Add “Details” tab in bottom panel
- Implement Describe/BatchGetBuilds detail transform: phases, env, artifacts
- Render human-readable details with JSON copy action
- Quick win: Open a build and read its phases/status clearly

## Week 5: Log Tail (CloudWatch Logs)
- From selected build, resolve log group/stream
- Implement a tailer:
	- Poll GetLogEvents with nextToken/nextStartTimestamp
	- Append to ring buffer; render in Logs tab
	- Handle backoff and cancellation
- Quick win: Live log lines streaming in terminal when build runs

## Week 6: Auto-Refresh and Backoff Polish
- Add timed auto-refresh loops with debounce for projects/builds
- Surface throttling/backoff in status bar and command log
- Add “Stop tail” and auto-stop when build completes
- Quick win: App keeps itself fresh without spamming AWS; smooth UX

## Week 7: Error UX + SSO Flow
- Detect expired creds/SSO; add “l” to open device auth flow (browser prompt text)
- Better toasts for common errors; retry guidance
- Quick win: Clear messages when creds expire; one key to fix

## Week 8: Quality, Tests, and Config
- Add unit tests for reducers and AWS wrappers (fakes)
- Add app config: keymap overrides, refresh intervals
- Persist minimal user preferences
- Quick win: Run go test and see green; tweak keybindings via config

## Stretch (Weeks 9–10): UX Refinements
- Copy commands (ARNs/IDs) to clipboard (use xclip/pbccopy when available)
- Build search by sourceVersion (commit SHA) if present
- Color themes and compact/expanded layouts

Each week should end with a visible, demo-able improvement on real AWS data.
