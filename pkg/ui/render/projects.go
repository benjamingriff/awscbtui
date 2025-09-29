package render

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func RenderProjects(v *gocui.View, s *state.AppState) {
	v.Clear()
	v.Title = "[1]-Projects"
	v.Wrap = true

	v.FrameColor = gocui.ColorGreen
	v.FgColor = gocui.ColorDefault
	v.SelBgColor = gocui.ColorGreen

	if s.UI.FocusedView == "projects" {
		v.Highlight = true
	} else {
		v.Highlight = false
	}

	for i, p := range s.Data.Projects {
		prefix := "  "
		if i == s.UI.SelectedProject {
			prefix = "> "
		}
		fmt.Fprintf(v, "%s%s  (%d builds)\n", prefix, p.Name)
	}
}

