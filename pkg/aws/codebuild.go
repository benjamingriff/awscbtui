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

func (c *CodeBuildClient) ListBuildsForProject(ctx context.Context, projectName string) ([]state.Build, error) {
	var builds []state.Build

	resp, err := c.cb.ListBuildsForProject(ctx, &codebuild.ListBuildsForProjectInput{
		ProjectName: &projectName,
	})
	if err != nil {
		return nil, err
	}
	for _, name := range resp.Ids {
		builds = append(builds, state.Build{ID: name})
	}
	return builds, nil
}

