package state

import (
	"time"
	"github.com/benjamingriff/awscbtui/pkg/aws"
)

type AppState struct {
	Session SessionState
	UI      UIState
	Data    DataState
	// Jobs    JobsState
	// CmdLog  CommandLog
	// Meta    MetaState
}

type SessionState struct {
	Profile   string
	Region    string
	Identity  string // ARN or display name
	// SSOStatus SSOStatus // enum: OK | EXPIRED | REQUIRES_LOGIN
	ExpiresAt time.Time
}


type UIState struct {
	FocusedView View
	SelectedProjectIdx int    // index in Projects slice
	SelectedBuildIdx   int    // index in Builds for selected project
	FilterText      string
	ShowHelp        bool
	// Toasts          []Toast // short transient messages
	// ScrollOrigins   map[string]image.Point // preserve origins per view (optional)
}


type DataState struct {
	Projects []aws.Project
	Builds map[string][]aws.Build
	BuildPhase map[string]aws.BuildPhase // keyed by buildID
	// LogTail map[string]*RingBuffer // keyed by buildID
}

type View string

const (
	ViewStatus View = "status"
	ViewProjects View = "projects"
	ViewBuilds View = "builds"
	ViewLogs View = "logs"
)
