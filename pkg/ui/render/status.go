package render

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func RenderStatus(v *gocui.View, s *state.AppState, focusedName string) {
	v.FrameRunes = []rune{'─', '│', '╭', '╮', '╰', '╯'}
	v.Title = "[0]-Status"
	v.Wrap = false

	v.FrameColor = gocui.ColorDefault
	v.FgColor = gocui.ColorDefault
	v.SelBgColor = gocui.ColorGreen

	isFocused := s.UI.FocusedView == state.ViewStatus
	v.Highlight = isFocused
	if isFocused {
		v.FrameColor = gocui.ColorGreen
	}

	if s.Session.Profile != "" {
		fmt.Fprintln(v, "Profile:", s.Session.Profile)
	} else {
		fmt.Fprintln(v, "Profile:")
	}

	fmt.Fprintln(v, "Current View:", focusedName)
}
