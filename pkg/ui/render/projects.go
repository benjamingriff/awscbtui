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

	v.Highlight = false
	prefix := "• "
	for i, p := range s.Data.Projects {
		if s.UI.FocusedView == "projects" && i == s.UI.SelectedProjectIdx {
			v.Highlight = true
		}
		fmt.Fprintf(v, "%s%s  (%d builds)\n", prefix, p.Name)
	}
}

