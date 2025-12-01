package aws

import (
	"context"

	sdkaws "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	cbtypes "github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

type CodeBuildClient struct {
  cb *codebuild.Client
}

func NewCodeBuildClient(cfg sdkaws.Config) *CodeBuildClient {
  return &CodeBuildClient{
    cb: codebuild.NewFromConfig(cfg),
  }
}

func (c *CodeBuildClient) ListProjects(ctx context.Context) ([]state.Project, error) {
	var projects []state.Project
	var next *string

	for {
		resp, err := c.cb.ListProjects(ctx, &codebuild.ListProjectsInput{
			NextToken: next,
			SortBy:    cbtypes.ProjectSortByTypeLastModifiedTime,
			SortOrder: cbtypes.SortOrderTypeDescending,
		})
		if err != nil {
			return nil, err
		}
		for _, name := range resp.Projects {
			projects = append(projects, state.Project{Name: name})
		}
		if resp.NextToken == nil {
			break
		}
		next = resp.NextToken
	}
	return projects, nil
}

func (c *CodeBuildClient) ListBuildsForProject(ctx context.Context, projectName string) ([]state.BuildId, error) {
	var buildIds []state.BuildId

	resp, err := c.cb.ListBuildsForProject(ctx, &codebuild.ListBuildsForProjectInput{
		ProjectName: &projectName,
	})
	if err != nil {
		return nil, err
	}
	for _, name := range resp.Ids {
		buildIds = append(buildIds, state.BuildId{ID: name})
	}
	return buildIds, nil
}

func (c *CodeBuildClient) BatchGetBuilds(ctx context.Context, buildIds []state.BuildId) ([]state.Build, error) {
	var ids []string
	var builds []state.Build

	for _, id := range buildIds {
		ids = append(ids, id.ID)
	}

	resp, err := c.cb.BatchGetBuilds(ctx, &codebuild.BatchGetBuildsInput{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	for _, build := range resp.Builds {
		builds = append(builds, toStateBuild(build))
	}
	return builds, nil
}

func toStateBuild(b cbtypes.Build) state.Build {
  var logGroup, logStream string
  if b.Logs != nil && b.Logs.CloudWatchLogs != nil {
    logGroup = sdkaws.ToString(b.Logs.CloudWatchLogs.GroupName)
    logStream = sdkaws.ToString(b.Logs.CloudWatchLogs.StreamName)
  }

  return state.Build{
    ID:        sdkaws.ToString(b.Id),
    ARN:       sdkaws.ToString(b.Arn),
    Project:   sdkaws.ToString(b.ProjectName),
    Status:    string(b.BuildStatus),
    StartTime: sdkaws.ToTime(b.StartTime),
    EndTime:   sdkaws.ToTime(b.EndTime),
    Initiator: sdkaws.ToString(b.Initiator),
    SourceRev: sdkaws.ToString(b.ResolvedSourceVersion),
    LogGroup:  logGroup,
    LogStream: logStream,
  }
}
