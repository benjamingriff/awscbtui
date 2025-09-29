package ui

import (
	"github.com/awesome-gocui/gocui"
	"github.com/benjamingriff/awscbtui/pkg/state"
)

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func MoveViewForward(s *state.AppState)  {
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
