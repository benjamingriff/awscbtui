package aws

import (
  "context"
  "time"

  awsv2 "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/service/codebuild"
  cbtypes "github.com/aws/aws-sdk-go-v2/service/codebuild/types"
  "github.com/benjamingriff/awscbtui/pkg/state"
)

type CodeBuildClient struct {
  cb *codebuild.Client
}

func NewCodeBuildClient(cfg awsv2.Config) *CodeBuildClient {
  return &CodeBuildClient{
    cb: codebuild.NewFromConfig(cfg),
  }
}

func (c *CodeBuildClient) ListProjects(ctx context.Context, limit int) ([]state.Project, error) {
  var out []state.Project
  var next *string
  remaining := limit
  for {
    if remaining <= 0 {
      break
    }
    pageSize := int32(remaining)
    if pageSize <= 0 || pageSize > 100 {
      pageSize = 100 // API max
    }
    resp, err := c.cb.ListProjects(ctx, &codebuild.ListProjectsInput{
      NextToken: next,
      SortBy:    cbtypes.ProjectSortByTypeName, // or LastModifiedTime if preferred
      SortOrder: cbtypes.SortOrderTypeAscending,
      MaxResults: &pageSize,
    })
    if err != nil {
      return nil, err
    }
    for _, name := range resp.Projects {
      out = append(out, state.Project{Name: name})
    }
    remaining = limit - len(out)
    if resp.NextToken == nil || remaining <= 0 {
      break
    }
    next = resp.NextToken
  }
  return out, nil
}
