package state

import (
	"github.com/benjamingriff/awscbtui/pkg/aws"
)

type Message interface {
	MessageName() string
}

type SessionLoaded struct {
  SessionInfo aws.SessionInfo
}
func (SessionLoaded) MessageName() string { return "SessionLoaded" }

type SessionError struct{ Err error }
