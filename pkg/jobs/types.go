package jobs

import (
	"context"
	sdkaws "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

type SessionAPI interface {
  LoadSession(ctx context.Context) (state.SessionInfo, sdkaws.Config, error)
}

type CodeBuildAPI interface {
  ListProjects(ctx context.Context) ([]state.Project, error)
	ListBuildsForProject(ctx context.Context, projectName string) ([]state.Build, error)
}

type jobKey struct {
    Kind   string
    Scope  string
}

