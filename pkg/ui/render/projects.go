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
	for _, p := range s.Data.Projects {
		fmt.Fprintf(v, "%s  (%d builds)\n", p.Name, len(s.Data.Projects))
		// fmt.Fprintf(v, "%s%s  (%d builds)\n", prefix, p.Name)
	}

	if v.Highlight {
		v.FrameColor = gocui.ColorGreen
		idx := s.UI.SelectedProjectIdx
		_, _ = ensureVisible(v, idx, len(s.Data.Projects))
	} else {
		_ = v.SetCursor(0, 0)
	}
}
