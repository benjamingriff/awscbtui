package state

func Reduce(s *AppState, a any) {
    switch v := a.(type) {
    case FocusNext:
        s.UI.FocusedView = nextPanel(s.UI.FocusedView)
    case FocusPrev:
        s.UI.FocusedView = prevPanel(s.UI.FocusedView)
	}
}

func nextPanel(p Panel) Panel {
    switch p {
    case PanelProjects:
        return PanelBuilds
    case PanelBuilds:
        return PanelBottom
    default:
        return PanelProjects
    }
}

func prevPanel(p Panel) Panel {
    switch p {
    case PanelBottom:
        return PanelBuilds
    case PanelBuilds:
        return PanelProjects
    default:
        return PanelBottom
    }
}
