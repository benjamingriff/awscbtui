package state

import (
	"time"
)

func NewAppState() AppState {
	return AppState{
		Session: SessionState{
			Profile:   "",
			Region:    "",
			Identity:  "",
			ExpiresAt: time.Time{},
		},
		UI: UIState{
			FocusedView:        ViewStatus,
			FocusedProjectIdx: 0,
			FocusedBuildIdx:   0,
			FilterText:         "",
			ShowHelp:           false,
			Loading: false,
		},
		Data: DataState{
			Projects:   make([]Project, 0),
			Builds:     make(map[string][]Build),
			BuildPhase: make(map[string]BuildPhase),
			// LogTail: make(map[string]*RingBuffer),
		},
	}
}

type AppState struct {
	Session SessionState
	UI UIState
	Data DataState
	// Jobs    JobsState
	// CmdLog  CommandLog
	// Meta    MetaState
}

type SessionState struct {
	Profile string
	Region string
	Identity string // ARN or display name
	// SSOStatus SSOStatus // enum: OK | EXPIRED | REQUIRES_LOGIN
	ExpiresAt time.Time
}

type SSOStatus string

const (
	SSOStatusOK SSOStatus = "OK"
	SSOStatusExpired SSOStatus = "EXPIRED"
	SSOStatusRequiresLogin SSOStatus = "REQUIRES_LOGIN"
)

type UIState struct {
	FocusedView View
	FocusedProjectIdx int
	SelectedProjectIdx int
	FocusedBuildIdx int
	SelectedBuildIdx int
	FilterText string
	ShowHelp bool
	Loading bool
	// Toasts          []Toast // short transient messages
	// ScrollOrigins   map[string]image.Point // preserve origins per view (optional)
}

type DataState struct {
	Projects []Project
	Builds map[string][]Build
	BuildPhase map[string]BuildPhase // keyed by buildID
	// LogTail map[string]*RingBuffer // keyed by buildID
}

type View string

const (
	ViewStatus View = "status"
	ViewProjects View = "projects"
	ViewBuilds View = "builds"
	ViewLogs View = "logs"
	ViewHelp View = "help"
)
