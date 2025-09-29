package state

func Reduce(s *AppState, a any) {
    switch a.(type) {
    case FocusNext:
        s.UI.FocusedView = nextPanel(s.UI.FocusedView)
    case FocusPrev:
        s.UI.FocusedView = prevPanel(s.UI.FocusedView)
	}
}

func nextPanel(p View) View {
	switch p {
	case ViewProjects:
		return ViewBuilds
	case ViewBuilds:
		return ViewLogs
	case ViewLogs:
		return ViewStatus
	case ViewStatus:
		return ViewProjects
	default:
		return ViewProjects
	}
}

func prevPanel(p View) View {
    switch p {
    case ViewProjects:
        return ViewStatus
    case ViewBuilds:
        return ViewProjects
    case ViewLogs:
        return ViewBuilds
    case ViewStatus:
        return ViewLogs
    default:
        return ViewProjects
    }
}
