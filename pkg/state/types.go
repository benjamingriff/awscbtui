package state 

import (
	"time"
)

type Project struct {
	Name string
}

type Build struct {
	ID string
}

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
