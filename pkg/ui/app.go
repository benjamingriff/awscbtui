package ui

import (
	"log"
	"github.com/jroimartin/gocui"
)

type Config struct {
	Name string
}

func Run(cfg *Config) error {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(Layout)

	if err := g.SetKeybinding("", 'q', gocui.ModNone, Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	return nil
}
