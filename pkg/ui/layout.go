package ui

import (
	"github.com/jroimartin/gocui"
	"github.com/benjamingriff/awscbtui/pkg/ui/render"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

const (
	ViewStatus = "status"
	ViewPipelines = "pipelines"
	ViewBuilds = "builds"
	ViewLogs = "logs"
)

func Layout(g *gocui.Gui, s *state.AppState) error {
	maxX, maxY := g.Size()
	statusH := maxY/5
	logsH:= maxY/5
	buildsH := maxY - statusH - logsH
	split := maxX/3

	if v, err := g.SetView(ViewStatus, 0, 0, maxX-1, statusH-1); err != nil {
		if err != gocui.ErrUnknownView { return err }
		render.RenderStatus(v, s)
	}

	if v, err := g.SetView(ViewPipelines, 0, statusH, split-1, buildsH+statusH-1); err != nil {
		if err != gocui.ErrUnknownView { return err }
		render.RenderProjects(v, s)
	}


	if v, err := g.SetView(ViewBuilds, split, statusH, maxX-1, buildsH+statusH-1); err != nil {
		if err != gocui.ErrUnknownView { return err }
		render.RenderBuilds(v, s)
	}


	if v, err := g.SetView(ViewLogs, 0, statusH+buildsH, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView { return err }
		render.RenderLogs(v,s)
	}

	return nil
}

