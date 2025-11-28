package state

type Message interface {
	MessageName() string
}

type SessionLoaded struct {
  SessionInfo SessionInfo
}
func (SessionLoaded) MessageName() string { return "SessionLoaded" }

type SessionError struct{ Err error }

type ProjectsLoaded struct {
  Projects []Project
}
func (ProjectsLoaded) MessageName() string { return "ProjectsLoaded" }
