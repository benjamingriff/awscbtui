package aws

import (
	"time"
)

type SessionInfo struct {
  Profile   string
  Region    string
  AccountID string
  ARN       string
  ExpiresAt *time.Time
  ErrorHint string
}

type Project struct {
	Name string
}

type Build struct {
	ID string
}

type BuildPhase struct {
	Name string
}
