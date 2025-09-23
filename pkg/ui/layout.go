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
		v.Title = "Status"
		v.Wrap = true
	}

	if v, err := g.SetView(ViewPipelines, 0, statusH, split-1, buildsH+statusH-1); err != nil {
		if err != gocui.ErrUnknownView { return err }
		v.Title = "Projects"
		v.Wrap = true
	}


	if v, err := g.SetView(ViewBuilds, split, statusH, maxX-1, buildsH+statusH-1); err != nil {
		if err != gocui.ErrUnknownView { return err }
		v.Title = "Builds"
		v.Wrap = true
	}


	if v, err := g.SetView(ViewLogs, 0, statusH+buildsH, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView { return err }
		v.Title = "Logs"
		v.Wrap = true
	}

	if v := mustView(g, ViewStatus); v != nil { render.RenderStatus(v, s) }
	if v := mustView(g, ViewPipelines); v != nil { render.RenderProjects(v, s) }
	if v := mustView(g, ViewBuilds); v != nil { render.RenderBuilds(v, s) }
	if v := mustView(g, ViewLogs); v != nil { render.RenderLogs(v, s) }
	return nil
}

func mustView(g *gocui.Gui, name string) *gocui.View {
	v, err := g.View(name) 
	if err != nil { return nil } 
	return v
}
