package ui

import (
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

type Binding struct {
	Rune rune
	Key gocui.Key
	Mod gocui.Modifier
	Handler func(*gocui.Gui, *gocui.View) error
	Help string
}

type Keymap struct {
	Bindings []Binding
}

func intentHandler(app *App, in state.Intent) func(*gocui.Gui, *gocui.View) error {
	return func(*gocui.Gui, *gocui.View) error {
		app.emit(in)
		return nil
	}
}

func KeymapDefault(app *App) Keymap {
  return Keymap{Bindings: []Binding{
		{Rune: 'h', Mod: gocui.ModNone, Handler: intentHandler(app, state.ViewPrev{}), Help: "prev"},
    {Rune: 'j', Mod: gocui.ModNone, Handler: intentHandler(app, state.IdxNext{}), Help: "down"},
    {Rune: 'k', Mod: gocui.ModNone, Handler: intentHandler(app, state.IdxPrev{}), Help: "up"},
		{Rune: 'l', Mod: gocui.ModNone, Handler: intentHandler(app, state.ViewNext{}), Help: "next"},
    {Rune: 'q', Mod: gocui.ModNone, Handler: intentHandler(app, state.Quit{}) , Help: "quit"},
    {Rune: '?', Mod: gocui.ModNone, Handler: intentHandler(app, state.RenderHelp{}) , Help: "help"},
  }}
}


func KeymapSpecial(app *App) Keymap {
  return Keymap{Bindings: []Binding{
    {Key: gocui.KeySpace, Mod: gocui.ModNone, Handler: intentHandler(app, state.RenderHelp{}) , Help: "select"},
  }}
}

func bindKeymaps(g *gocui.Gui, view string, keymap Keymap) error {
  for _, b := range keymap.Bindings {
    switch {
    case b.Rune != 0:
      if err := g.SetKeybinding(view, b.Rune, b.Mod, b.Handler); err != nil {
        return err
      }
    default:
      if err := g.SetKeybinding(view, b.Key, b.Mod, b.Handler); err != nil {
        return err
      }
    }
  }
  return nil
}
