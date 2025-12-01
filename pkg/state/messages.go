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

type BuildIdsLoaded struct {
	ProjectName string
  BuildIds []BuildId
}
func (BuildIdsLoaded) MessageName() string { return "BuildIdsLoaded" }

type BuildsLoaded struct {
	ProjectName string
  Builds []Build
}
func (BuildsLoaded) MessageName() string { return "BuildsLoaded" }
