package ui

import (
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/ui/render"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func Layout(g *gocui.Gui, s *state.AppState) error {
	maxX, maxY := g.Size()
	statusH := maxY/8
	logsH:= maxY/2
	buildsH := maxY - statusH - logsH
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

	if _, err := g.SetView("logs", 0, statusH+buildsH, maxX-1, maxY-1, 0); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	if v, err := g.View("logs"); err == nil {
		v.Clear()
		render.RenderLogs(v, s)
	}

	return nil
}
