package render

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func RenderStatus(v *gocui.View, s *state.AppState) {
	v.Clear()
	v.Title = "[0]-Status"
	v.Wrap = true

	v.FrameColor = gocui.ColorGreen
	v.FgColor = gocui.ColorMagenta

	fmt.Fprintln(v, "AWS accout status")
}
