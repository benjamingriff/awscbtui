package render

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func RenderLogs(v *gocui.View, s *state.AppState) {
	v.Clear()
	v.Title = "Logs"
	v.Wrap = true

	fmt.Fprintln(v, "This is where the logs live")
}
