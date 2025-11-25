package aws

import (
	"context"
	"os"
	"time"
	sdkaws "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

type SessionClient struct{}

func NewSessionClient() *SessionClient { return &SessionClient{} }

// func (s *SessionClient) LoadSession(ctx context.Context, profile, region string) (state.SessionState, error) {
func (s *SessionClient) LoadSession(ctx context.Context) (state.SessionInfo, error) {
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

  return state.SessionInfo{
    Profile:   profileFromConfig(),
    Region:    cfg.Region,
    AccountID: sdkaws.ToString(idOut.Account),
    ARN:       sdkaws.ToString(idOut.Arn),
    ExpiresAt: expPtr,
    ErrorHint: "",
  }, nil
}

func profileFromConfig() string {
  if p := os.Getenv("AWS_PROFILE"); p != "" {
    return p
  }
  return "default"
}
