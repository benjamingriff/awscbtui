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
			FocusedProjectIdx: -1,
			SelectedProjectIdx: -1,
			FocusedBuildIdx:   -1,
			SelectedBuildIdx:   -1,
			SelectedProjectName: "",
			SelectedBuildID:     "",
			FilterText:         "",
			ShowHelp:           false,
			Loading: false,
		},
		Data: DataState{
			Projects:   make([]Project, 0),
			BuildIds:     make(map[string][]BuildId),
			Builds:     make(map[string][]Build),
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
  FocusedView         View
  FocusedProjectIdx   int
  SelectedProjectIdx  int
  FocusedBuildIdx     int
  SelectedBuildIdx    int
  SelectedProjectName string
  SelectedBuildID     string
  FilterText          string
  ShowHelp            bool
  Loading             bool
	// Toasts          []Toast // short transient messages
	// ScrollOrigins   map[string]image.Point // preserve origins per view (optional)
}

type DataState struct {
	Projects []Project
	BuildIds map[string][]BuildId
	Builds map[string][]Build 
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


type Project struct {
	Name string
}

type BuildId struct {
	ID string
}

type Build struct {
	ID        string
	BuildNumber int64
	ARN       string
	Project   string
	Status    string
	StartTime time.Time
	EndTime   time.Time
	Initiator string
	SourceRev string
	LogGroup  string
	LogStream string
}

type StatusType struct {
	Status string
}

type Status string 

const ( 
	SUCCEEDED Status = "SUCCEEDED"
	FAILED Status = "FAILED"
	FAULT Status = "FAULT"
	IN_PROGRESS Status = "IN_PROGRESS"
	STOPPED Status = "STOPPED"
	TIMED_OUT Status = "TIMED_OUT"
)


type BuildPhase struct {
	Name string
}

type SessionInfo struct {
  Profile   string
  Region    string
  AccountID string
  ARN       string
  ExpiresAt *time.Time
  ErrorHint string
}
