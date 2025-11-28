package render

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func RenderProjects(v *gocui.View, s *state.AppState) {
	v.FrameRunes = []rune{'─', '│', '╭', '╮', '╰', '╯'}
	v.Title = "[1]-Projects"
	v.Wrap = false
	v.FrameColor = gocui.ColorDefault
	v.FgColor = gocui.ColorDefault
	v.SelBgColor = gocui.ColorGreen
	v.SelFgColor = gocui.ColorBlack

	v.Highlight = (s.UI.FocusedView == state.ViewProjects)

	v.Clear()
	for i, p := range s.Data.Projects {
		if s.UI.SelectedBuildIdx == i {
			fmt.Fprintf(v, " \x1b[32m*\x1b[0m %s \n", p.Name)
		} else {
			fmt.Fprintf(v, "   %s \n", p.Name)
		}
	}

	if v.Highlight {
		v.FrameColor = gocui.ColorGreen
		idx := s.UI.FocusedProjectIdx
		_, _ = ensureVisible(v, idx, len(s.Data.Projects))
	} else {
		_ = v.SetCursor(0, 0)
	}
}
