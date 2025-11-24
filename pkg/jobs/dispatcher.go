package jobs 

import (
	"os"
	"context"
	"time"
	sdkaws "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/benjamingriff/awscbtui/pkg/state"
	"github.com/benjamingriff/awscbtui/pkg/aws"
)

type Dispatcher struct {
	msgCh   chan<- state.Message
    inflight map[jobKey]context.CancelFunc
}

func NewDispatcher(
    msgCh chan<- state.Message,
) *Dispatcher {
    return &Dispatcher{
        msgCh:    msgCh,
        inflight:  make(map[jobKey]context.CancelFunc),
    }
}

func (d *Dispatcher) FetchProjects(ctx context.Context) error {
	return nil
}

func (d *Dispatcher) DispatchLoadSession(ctx context.Context) {
  go d.loadSessionWorker(ctx)
}

func (d *Dispatcher) loadSessionWorker(ctx context.Context) {
  opts := []func(*config.LoadOptions) error{}

  // if profile != "" {
  //   opts = append(opts, config.WithSharedConfigProfile(profile))
  // }
  // if region != "" {
  //   opts = append(opts, config.WithRegion(region))
  // }

  cfg, _ := config.LoadDefaultConfig(ctx, opts...)
  // if err != nil {
  //   d.out <- state.JobError{
  //     Key:      jobSession,
  //     ErrKind:  classifyAWSError(err),
  //     Err:      err,
  //     UserHint: "Failed to load AWS config. Check profile/region or run `aws configure`/SSO login.",
  //   }
  //   return
  // }

  stscli := sts.NewFromConfig(cfg)
  idOut, _ := stscli.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
  // if err != nil {
  //   d.out <- state.JobError{
  //     Key:      jobSession,
  //     ErrKind:  classifyAWSError(err),
  //     Err:      err,
  //     UserHint: ssoHintForError(err),
  //   }
  //   return
  // }

  var expPtr *time.Time
  if cred, credErr := cfg.Credentials.Retrieve(ctx); credErr == nil && cred.CanExpire {
    t := cred.Expires.UTC()
    expPtr = &t
  }

  info := aws.SessionInfo{
    Profile:   profileFromConfig(),
    Region:    cfg.Region,
    AccountID: sdkaws.ToString(idOut.Account),
    ARN:       sdkaws.ToString(idOut.Arn),
    ExpiresAt: expPtr,
    ErrorHint: "",
  }

  d.msgCh <- state.SessionLoaded{SessionInfo: info}
}

func profileFromConfig() string {
  if p := os.Getenv("AWS_PROFILE"); p != "" {
    return p
  }
  return "default"
}
