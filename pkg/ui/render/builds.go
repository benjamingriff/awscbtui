package render

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func RenderBuilds(v *gocui.View, s *state.AppState) {
	v.Clear()
	fmt.Fprintln(v, "Hi I'm the builds!")
	// builds := s.BuildsByProject[projectName]
	// for i, b := range builds {
	// 	sel := " "
	// 	if i == s.UI.SelectedBuildIndex { sel = ">" }
	// 	fmt.Fprintf(v, "%s %s %-7s %s\n", sel, b.BuildID, b.Status, humanDuration(b.Duration))
	// }
}
