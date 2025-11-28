package render

import (
	"fmt"
	"strings"
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func RenderCmds(v *gocui.View, s *state.AppState) {
	v.Clear()
	v.Frame = false
	v.Wrap = true
	v.FgColor = gocui.ColorBlue
	
	cmds := []string{
		"Do Something: <space>",
		"Keybindings: ?",
	}

	cmdString := strings.Join(cmds, " | ")

	fmt.Fprintln(v, cmdString)
}
