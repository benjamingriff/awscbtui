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

	if err := g.SetKeybinding(
		"",
		'l',
		gocui.ModNone,
		func(gg *gocui.Gui, v *gocui.View) error {
			MoveViewForwards(s)
			gg.Update(func(*gocui.Gui) error { return nil })
			return nil
		},
		); err != nil {
		return err
	}

	if err := g.SetKeybinding(
		"",
		'h',
		gocui.ModNone,
		func(gg *gocui.Gui, v *gocui.View) error {
			MoveViewBackwards(s)
			gg.Update(func(*gocui.Gui) error { return nil })
			return nil
		},
		); err != nil {
		return err
	}

	if err := g.SetKeybinding(
		"",
		'k',
		gocui.ModNone,
		func(gg *gocui.Gui, v *gocui.View) error {
			MoveIdxForwards(s)
			gg.Update(func(*gocui.Gui) error { return nil })
			return nil
		},
		); err != nil {
		return err
	}

	if err := g.SetKeybinding(
		"",
		'j',
		gocui.ModNone,
		func(gg *gocui.Gui, v *gocui.View) error {
			MoveIdxBackwards(s)
			gg.Update(func(*gocui.Gui) error { return nil })
			return nil
		},
		); err != nil {
		return err
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	return nil
}
