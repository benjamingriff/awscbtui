package state 

import (
	"time"
)

type Project struct {
	Name string
}

type BuildId struct {
	ID string
}

type Build struct {
	ID        string
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
