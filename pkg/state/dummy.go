package state

import (
	"time"
	"github.com/benjamingriff/awscbtui/pkg/aws"
)

func DummyState() *AppState {
    return &AppState{
        Session: SessionState{
            Profile:   "default",
            Region:    "us-east-1",
            Identity:  "arn:aws:iam::123456789012:user/test",
            ExpiresAt: time.Now().Add(1 * time.Hour),
        },
        UI: UIState{
            FocusedPanel:    "projects",
            ActiveTab:       "logs",
            SelectedProject: 0,
            SelectedBuild:   0,
            FilterText:      "",
            ShowHelp:        false,
        },
        Data: DataState{
            Projects: []aws.Project{
                {Name: "project-alpha"},
                {Name: "project-beta"},
                {Name: "project-gamma"},
            },
            Builds: map[string][]aws.Build{
                "project-alpha": {
                    {ID: "build-1"},
                    {ID: "build-2"},
                },
                "project-beta": {
                    {ID: "build-11"},
                },
            },
            BuildPhase: map[string]aws.BuildPhase{
                "build-1": {Name: "BUILD"},
            },
        },
    }
}
