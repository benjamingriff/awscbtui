package ui

import (
	"github.com/awesome-gocui/gocui"
)

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
