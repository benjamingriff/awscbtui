package ui

import (
	"github.com/jroimartin/gocui"
)

func QuitAndRestart(g *gocui.Gui, v *gocui.View) error {
	gocui.ErrQuit()
	return Run()
	
}

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
