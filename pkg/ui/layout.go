package ui

import (
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/ui/render"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func Layout(g *gocui.Gui, app *App) error {
	maxX, maxY := g.Size()
	cmdsH := 2
	useY := maxY - cmdsH
	useYF := float64(useY)
	statusH := useY/8

	var logsH int
	if app.state.UI.FocusedView == state.ViewLogs {
		logsH = int(useYF * 0.75) 
	} else {
		logsH = int(useYF / 5)
	}

	buildsH := useY - statusH - logsH
	split := maxX/5

	focusedName := getCurrentViewName(g)

	if _, err := g.SetView("status", 0, 0, split-1, statusH-1, 0); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	if v, err := g.View("status"); err == nil {
		if app.state.UI.FocusedView == state.ViewStatus {
			_, _ = g.SetCurrentView("status")
		}
		v.Clear()
		render.RenderStatus(v, &app.state, focusedName)
	}

	if _, err := g.SetView("projects", 0, statusH, split-1, buildsH+statusH-1, 0); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	if v, err := g.View("projects"); err == nil {
		if app.state.UI.FocusedView == state.ViewProjects {
			_, _ = g.SetCurrentView("projects")
		}
		if err := bindKeymaps(g, "projects", KeymapProjects(app)); err != nil {
			return err
		}
		v.Clear()
		render.RenderProjects(v, &app.state)
	}

	if _, err := g.SetView("builds", split, 0, maxX-1, buildsH+statusH-1, 0); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	if v, err := g.View("builds"); err == nil {
		if app.state.UI.FocusedView == state.ViewBuilds {
			_, _ = g.SetCurrentView("builds")
		}
		if err := bindKeymaps(g, "builds", KeymapBuilds(app)); err != nil {
			return err
		}
		v.Clear()
		render.RenderBuilds(v, &app.state)
	}

	if _, err := g.SetView("logs", 0, statusH+buildsH, maxX-1, maxY-cmdsH, 0); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	if v, err := g.View("logs"); err == nil {
		if app.state.UI.FocusedView == state.ViewLogs {
			_, _ = g.SetCurrentView("logs")
		}
		v.Clear()
		render.RenderLogs(v, &app.state)
	}


	if app.state.UI.ShowHelp {
		if _, err := g.SetView("help", maxX/4, maxY/6, (maxX/4)*3, (maxY/6)*5, 0); err != nil && err != gocui.ErrUnknownView {
			return err
		}
		if v, err := g.View("help"); err == nil {
			if app.state.UI.ShowHelp {
				_, _ = g.SetCurrentView("help")
			}
			v.Clear()
			render.RenderHelp(v, &app.state)
		}
	} else {
		if v, _ := g.View("help"); v != nil {
			if err := g.DeleteView("help"); err != nil {
				return err
			}
		}
	}

	if _, err := g.SetView("cmds", 0, statusH+buildsH+logsH, maxX-1, maxY, 0); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	if v, err := g.View("cmds"); err == nil {
		v.Clear()
		if app.state.UI.FocusedView == state.ViewStatus {
			render.RenderCmds(v, &app.state, CmdsStatus())
		} else if app.state.UI.FocusedView == state.ViewProjects {
			render.RenderCmds(v, &app.state, CmdsProjects())
		} else if app.state.UI.FocusedView == state.ViewBuilds {
			render.RenderCmds(v, &app.state, CmdsBuilds())
		} else if app.state.UI.FocusedView == state.ViewLogs {
			render.RenderCmds(v, &app.state, CmdsLogs())
		} else if app.state.UI.FocusedView == state.ViewHelp {
			render.RenderCmds(v, &app.state, CmdsHelp())
		}
	}

	return nil
}

func getCurrentViewName(g *gocui.Gui) string {
	cv := g.CurrentView()
	var focusedName string
	if cv != nil {
		focusedName = cv.Name()
	} else {
		focusedName = "" 
	}
	return focusedName
}

func CmdsStatus() state.CmdHints {
  return state.CmdHints{Cmds: []state.Cmd{
		{Text: "Back", Key: "h"},
		{Text: "Down", Key: "j"},
		{Text: "Up", Key: "k"},
		{Text: "Next", Key: "l"},
		{Text: "Enter", Key: "<space>"},
		{Text: "Keymaps", Key: "?"},
  }}
}

func CmdsProjects() state.CmdHints {
  return state.CmdHints{Cmds: []state.Cmd{
		{Text: "Enter", Key: "<space>"},
		{Text: "Keymaps", Key: "?"},
  }}
}

func CmdsBuilds() state.CmdHints {
	return state.CmdHints{Cmds: []state.Cmd{
		{Text: "Enter", Key: "<space>"},
		{Text: "Keymaps", Key: "?"},
	}}
}

func CmdsLogs() state.CmdHints {
	return state.CmdHints{Cmds: []state.Cmd{
		{Text: "Enter", Key: "<space>"},
		{Text: "Keymaps", Key: "?"},
	}}
}

func CmdsHelp() state.CmdHints {
	return state.CmdHints{Cmds: []state.Cmd{
		{Text: "Exit", Key: "q"},
	}}
}
