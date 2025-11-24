package ui

import (
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/ui/render"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func Layout(g *gocui.Gui, s *state.AppState) error {
	maxX, maxY := g.Size()
	cmdsH := 2
  useY := maxY - cmdsH
	useYF := float64(useY)
	statusH := useY/8

	var logsH int
	if s.UI.FocusedView == state.ViewLogs {
		logsH = int(useYF * 0.75) 
	} else {
		logsH = int(useYF / 5)
	}

	buildsH := useY - statusH - logsH
	split := maxX/3

	if _, err := g.SetView("status", 0, 0, split-1, statusH-1, 0); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	if v, err := g.View("status"); err == nil {
		v.Clear()
		render.RenderStatus(v, s)
	}

	if _, err := g.SetView("projects", 0, statusH, split-1, buildsH+statusH-1, 0); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	if v, err := g.View("projects"); err == nil {
		v.Clear()
		render.RenderProjects(v, s)
	}

	if _, err := g.SetView("builds", split, 0, maxX-1, buildsH+statusH-1, 0); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	if v, err := g.View("builds"); err == nil {
		v.Clear()
		render.RenderBuilds(v, s)
	}

	if _, err := g.SetView("logs", 0, statusH+buildsH, maxX-1, maxY-cmdsH, 0); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	if v, err := g.View("logs"); err == nil {
		v.Clear()
		render.RenderLogs(v, s)
	}

	if _, err := g.SetView("cmds", 0, statusH+buildsH+logsH, maxX-1, maxY, 0); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	if v, err := g.View("cmds"); err == nil {
		v.Clear()
		render.RenderCmds(v, s)
	}

	if s.UI.ShowHelp {
		if _, err := g.SetView("help", maxX/4, maxY/6, (maxX/4)*3, (maxY/6)*5, 0); err != nil && err != gocui.ErrUnknownView {
			return err
		}
		if v, err := g.View("help"); err == nil {
			v.Clear()
			render.RenderHelp(v, s)
		}
	} else {
		if v, _ := g.View("help"); v != nil {
			if err := g.DeleteView("help"); err != nil {
				return err
			}
		}
	}

	return nil
}
