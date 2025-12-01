package jobs

import (
	"context"

	"github.com/benjamingriff/awscbtui/pkg/aws"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

type Dispatcher struct {
	msgCh   chan<- state.Message
	inflight map[jobKey]context.CancelFunc
	sesh SessionAPI
	cb CodeBuildAPI
}

func NewDispatcher(
    msgCh chan<- state.Message,
) *Dispatcher {
    return &Dispatcher{
        msgCh:    msgCh,
        inflight:  make(map[jobKey]context.CancelFunc),
				sesh: aws.NewSessionClient(),
    }
}

func (d *Dispatcher) DispatchLoadSession(ctx context.Context) {
  go func() {
    info, cfg, _ := d.sesh.LoadSession(ctx)
		d.cb = aws.NewCodeBuildClient(cfg)
    d.msgCh <- state.SessionLoaded{SessionInfo: info}
  }()
}

func (d *Dispatcher) FetchProjects(ctx context.Context) {
	go func() {
		projects, _ := d.cb.ListProjects(ctx)
    d.msgCh <- state.ProjectsLoaded{Projects: projects}
	}()
}

func (d *Dispatcher) LoadProjectsBuilds(ctx context.Context, projectName string) {
	go func() {
		builds, _ := d.cb.ListBuildsForProject(ctx, projectName)
		d.msgCh <- state.BuildsLoaded{ProjectName: projectName, Builds: builds}
	}()
}
