package ui

import (
	"log"
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func Run() error {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

 	s := state.DummyState()

	g.SetManagerFunc(func(g *gocui.Gui) error {
		return Layout(g, s)
	})

	if err := g.SetKeybinding("", 'q', gocui.ModNone, Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	return nil
}
