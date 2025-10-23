package ui

import (
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/ui/render"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func Quit(g *gocui.Gui, v *gocui.View, s *state.AppState) error {
	if s.UI.ShowHelp {
		if err := HideHelp(g, s); err != nil {
			return err
		}
		g.Update(func(*gocui.Gui) error { return nil })
		return nil
	}
	return gocui.ErrQuit
}

func MoveViewForwards(s *state.AppState)  {
	switch s.UI.FocusedView {
	case "status":
		s.UI.FocusedView = state.ViewProjects
	case "projects":
		s.UI.FocusedView = state. ViewBuilds
	case "builds":
		s.UI.FocusedView = state.ViewLogs 
	case "logs":
		s.UI.FocusedView = state.ViewStatus
	}
}
func MoveViewBackwards(s *state.AppState)  {
	switch s.UI.FocusedView {
	case "status":
		s.UI.FocusedView = state.ViewLogs
	case "projects":
		s.UI.FocusedView = state.ViewStatus
	case "builds":
		s.UI.FocusedView = state.ViewProjects
	case "logs":
		s.UI.FocusedView = state.ViewBuilds
	}
}

func MoveIdxForwards(s *state.AppState) {
	switch s.UI.FocusedView {
	case "projects":
		if len(s.Data.Projects) == 0 {
			return
		}
		if s.UI.SelectedProjectIdx == len(s.Data.Projects)-1 {
			s.UI.SelectedProjectIdx = 0
		} else {
			s.UI.SelectedProjectIdx++
		}
	case "builds":
		if len(s.Data.Builds) == 0 {
			return
		}
		if s.UI.SelectedBuildIdx == len(s.Data.Builds)-1 {
			s.UI.SelectedBuildIdx = 0
		} else {
			s.UI.SelectedBuildIdx++
		}
	}
}

func MoveIdxBackwards(s *state.AppState) {
	switch s.UI.FocusedView {
	case "projects":
		if len(s.Data.Projects) == 0 {
			return
		}
		if s.UI.SelectedProjectIdx == 0 {
			s.UI.SelectedProjectIdx = len(s.Data.Projects) - 1
		} else {
			s.UI.SelectedProjectIdx--
		}
	case "builds":
		if len(s.Data.Builds) == 0 {
			return
		}
		if s.UI.SelectedBuildIdx == 0 {
			s.UI.SelectedBuildIdx = len(s.Data.Builds) - 1
		} else {
			s.UI.SelectedBuildIdx--
		}
	}
}

func RenderHelp(g *gocui.Gui, s *state.AppState) error {
	maxX, maxY := g.Size()
	if _, err := g.SetView("help", maxX/4, maxY/6, (maxX/4)*3, (maxY/6)*5, 0); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	if v, err := g.View("help"); err == nil {
		v.Clear()
		render.RenderHelp(v, s)
	}
	s.UI.ShowHelp = true
	return nil
}

func HideHelp(g *gocui.Gui, s *state.AppState) error {
    if v, _ := g.View("help"); v != nil {
        if err := g.DeleteView("help"); err != nil {
            return err
        }
    }
    s.UI.ShowHelp = false
    return nil
}
