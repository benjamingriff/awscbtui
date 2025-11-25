package jobs

import (
	"context"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

type SessionAPI interface {
  // LoadSession(ctx context.Context, profile, region string) (state.SessionState, error)
  LoadSession(ctx context.Context) (state.SessionInfo, error)
}

type jobKey struct {
    Kind   string
    Scope  string
}

