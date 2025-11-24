package ui

import (
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func bindKeys(g *gocui.Gui, app *App) error {
	if err := g.SetKeybinding("", 'q', gocui.ModNone, func(*gocui.Gui, *gocui.View) error {
		app.emit(state.Quit{})
		return nil
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'l', gocui.ModNone, func(*gocui.Gui, *gocui.View) error {
		app.emit(state.ViewNext{})
		return nil
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'h', gocui.ModNone, func(*gocui.Gui, *gocui.View) error {
		app.emit(state.ViewPrev{})
		return nil
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'j', gocui.ModNone, func(*gocui.Gui, *gocui.View) error {
		app.emit(state.IdxNext{})
		return nil
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'k', gocui.ModNone, func(*gocui.Gui, *gocui.View) error {
		app.emit(state.IdxPrev{})
		return nil
	}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", '?', gocui.ModNone, func(*gocui.Gui, *gocui.View) error {
		app.emit(state.RenderHelp{})
		return nil
	}); err != nil {
		return err
	}
	// // refresh based on focus
	// if err := g.SetKeybinding("", 'r', gocui.ModNone, func(*gocui.Gui, *gocui.View) error {
	// 	switch app.State.FocusedPanel {
	// 	case state.PanelProjects:
	// 		app.emit(state.RefreshProjects{})
	// 	case state.PanelBuilds:
	// 		if p := selectors.CurrentProject(app.State); p != "" {
	// 			app.emit(state.RefreshBuilds{Project: p})
	// 		}
	// 	}
	// 	return nil
	// }); err != nil {
	// 	return err
	// }

	return nil
}
