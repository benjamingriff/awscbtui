package render

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func RenderCmds(v *gocui.View, s *state.AppState, cmds state.CmdHints) {
	v.Clear()
	v.Frame = false
	v.Wrap = true
	v.FgColor = gocui.ColorBlue
	
	cmdString := ""
	for _, cmd := range cmds.Cmds {
		cmdString += cmd.Text + ": " + cmd.Key + " | "
	}

	fmt.Fprintln(v, cmdString)
}
