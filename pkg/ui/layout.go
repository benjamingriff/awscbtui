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
	statusH := useY/8
	logsH:= useY/2
	buildsH := useY - statusH - logsH
	split := maxX/3

	if _, err := g.SetView("status", 0, 0, maxX-1, statusH-1, 0); err != nil && err != gocui.ErrUnknownView {
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

	if _, err := g.SetView("builds", split, statusH, maxX-1, buildsH+statusH-1, 0); err != nil && err != gocui.ErrUnknownView {
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

	return nil
}
