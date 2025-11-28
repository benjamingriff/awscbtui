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

func (d *Dispatcher) FetchProjects(ctx context.Context) error {
	return nil
}

func (d *Dispatcher) DispatchLoadSession(ctx context.Context) {
  go func() {
    info, cfg, _ := d.sesh.LoadSession(ctx)
    // if err != nil {
    //   d.msgCh <- state.JobError{
    //     Key:      "session:load",
    //     ErrKind:  state.ErrKindAuth, // or classify later
    //     Err:      err,
    //     UserHint: "Failed to load AWS session. Set AWS_PROFILE or choose a profile.",
    //   }
    //   return
    // }
    d.msgCh <- state.SessionLoaded{SessionInfo: info}

		d.cb = aws.NewCodeBuildClient(cfg)
		projects, _ := d.cb.ListProjects(ctx)
    d.msgCh <- state.ProjectsLoaded{Projects: projects}
  }()
}
