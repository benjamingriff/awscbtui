package render

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func RenderStatus(v *gocui.View, s *state.AppState) {
	v.Clear()
	v.Title = "Status"
	v.Wrap = true

	v.FgColor = gocui.ColorMagenta
	v.SelFgColor = gocui.ColorCyan

	fmt.Fprintln(v, "AWS accout status")
}
