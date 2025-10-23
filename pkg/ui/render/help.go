package render

import (
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func RenderHelp(v *gocui.View, s *state.AppState) {
	v.FrameRunes = []rune{'─', '│', '╭', '╮', '╰', '╯'}
	v.Title = "Keymaps"
	v.Wrap = true

	v.FrameColor = gocui.ColorGreen
	v.FgColor = gocui.ColorDefault
}
