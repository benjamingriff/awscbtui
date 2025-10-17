package render

import (
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func RenderBuilds(v *gocui.View, s *state.AppState) {
	v.Clear()
	v.Title = "[2]-Builds"
	v.Wrap = true

	v.FrameColor = gocui.ColorGreen
	v.FgColor = gocui.ColorDefault
	v.SelBgColor = gocui.ColorGreen

	if s.UI.FocusedView == "builds" {
		v.Highlight = true
	} else {
		v.Highlight = false
	}

	// for i, b := range s.Data.Builds {
	// 	sel := " "
	// 	if i == s.UI.SelectedBuildIdx {
	// 		sel = "> "
	// 	}
	// 	fmt.Fprintf(v, "%s %s %-7s %s\n", sel, b.BuildID, b.Status, humanDuration(b.Duration))
	// }
}
