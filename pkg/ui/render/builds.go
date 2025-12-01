package render

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func RenderBuilds(v *gocui.View, s *state.AppState) {
	v.FrameRunes = []rune{'─', '│', '╭', '╮', '╰', '╯'}
	v.Title = "[2]-Builds"
	v.Wrap = false
	v.Clear()

	v.FrameColor = gocui.ColorDefault
	v.FgColor = gocui.ColorDefault
	v.SelBgColor = gocui.ColorGreen

	if s.UI.FocusedView == "builds" {
		v.Highlight = true
		v.FrameColor = gocui.ColorGreen
	} else {
		v.Highlight = false
	}

	for _, builds := range s.Data.Builds {
		for _, b := range builds {
			fmt.Fprintf(v, "  Build ID: %s \n", b.ID)
		}
	}
}
