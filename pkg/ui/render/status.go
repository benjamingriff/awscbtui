package render

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func RenderStatus(v *gocui.View, s *state.AppState) {
	v.Clear()
	fmt.Fprintln(v, "AWS accout status")
}
