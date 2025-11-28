package render

import (
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func RenderLogs(v *gocui.View, s *state.AppState) {
	v.FrameRunes = []rune{'─', '│', '╭', '╮', '╰', '╯'}
	v.Clear()
	v.Title = "[3]-Logs"
	v.Wrap = true

	v.FrameColor = gocui.ColorDefault
	v.FgColor = gocui.ColorDefault
	v.SelBgColor = gocui.ColorGreen
	if s.UI.FocusedView == "logs" {
		v.Highlight = true
		v.FrameColor = gocui.ColorGreen
	} else {
		v.Highlight = false
	}
}
