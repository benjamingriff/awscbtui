Beginner Go tips relevant to this structure
- Keep the UI thread single-threaded: Only mutate AppState from one goroutine (usually the main UI loop). Use channels to pass results from workers to the UI.
- Use contexts everywhere: func(ctx context.Context, …) and cancel them when switching profile/region or when leaving a screen. This prevents runaway goroutines.
- Build small, vertical slices: Add one feature end-to-end (intent -> job -> state -> render), commit, then the next. Don’t build all layers at once.
- Interfaces at boundaries: Start concrete. Introduce small interfaces only where you need to swap implementations (e.g., AWS client vs fake).
- Error handling: Always check and return errors. Surface them in the command log and as a toast. Don’t swallow errors silently.
- Formatting and linting: Use go fmt (built-in) and consider golangci-lint. It’ll catch common mistakes early.
- Dependency management: Don’t commit unused packages. Run go mod tidy regularly.
- Testing: Focus on reducer tests first—they’re easiest and give high confidence. Add integration tests later.

Common pitfalls to avoid
- Doing network calls inside render functions. Render should be fast and pure.
- Updating gocui views from worker goroutines. Always send a message and update views on the UI/main goroutine.
- Forgetting pagination: CodeBuild and Logs APIs often paginate—write helpers that fetch all or stream progressively if needed.
- Tight polling loops: Respect rate limits. Use backoff and sensible intervals (1–2s for log tail, 15–60s for lists).
