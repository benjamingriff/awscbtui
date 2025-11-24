package state

import (
	"time"
	"github.com/benjamingriff/awscbtui/pkg/aws"
)
func DummyState() AppState {
	s := NewAppState()
	s.Data.Projects = []aws.Project{
		{Name: "proj-alpha"},
		{Name: "proj-beta"},
	}
	s.Data.Builds["proj-alpha"] = []aws.Build{
		{
			ID:        "alpha-001",
		},
		{
			ID:        "alpha-002",
		},
	}
	s.Data.Builds["proj-beta"] = []aws.Build{
		{
			ID:        "beta-001",
		},
		{
			ID:        "beta-002",
		},
	}
	return s
}

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
			SelectedProjectIdx: 0,
			SelectedBuildIdx:   0,
			FilterText:         "",
			ShowHelp:           false,
			Loading: false,
		},
		Data: DataState{
			Projects:   make([]aws.Project, 0),
			Builds:     make(map[string][]aws.Build),
			BuildPhase: make(map[string]aws.BuildPhase),
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
	SelectedProjectIdx int
	SelectedBuildIdx int
	FilterText string
	ShowHelp bool
	Loading bool
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
