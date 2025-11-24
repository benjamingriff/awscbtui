package state

func reduceQuit(s *AppState) []Effect {
	if s.UI.ShowHelp {
		s.UI.ShowHelp = false
		return nil
	}
	return []Effect{{Kind: CloseProgram}}
}

func reduceViewNext(s *AppState) []Effect  {
	switch s.UI.FocusedView {
	case "status":
		s.UI.FocusedView = ViewProjects
	case "projects":
		s.UI.FocusedView = ViewBuilds
	case "builds":
		s.UI.FocusedView = ViewLogs 
	case "logs":
		s.UI.FocusedView = ViewStatus
	}
	return nil
}

func reduceViewPrev(s *AppState) []Effect {
	switch s.UI.FocusedView {
	case "status":
		s.UI.FocusedView = ViewLogs
	case "projects":
		s.UI.FocusedView = ViewStatus
	case "builds":
		s.UI.FocusedView = ViewProjects
	case "logs":
		s.UI.FocusedView = ViewBuilds
	}
	return nil
}

func reduceIdxNext(s *AppState) []Effect {
	switch s.UI.FocusedView {
	case "projects":
		if len(s.Data.Projects) == 0 {
			return nil
		}
		if s.UI.SelectedProjectIdx == len(s.Data.Projects)-1 {
			s.UI.SelectedProjectIdx = 0
		} else {
			s.UI.SelectedProjectIdx++
		}
	case "builds":
		if len(s.Data.Builds) == 0 {
			return nil
		}
		if s.UI.SelectedBuildIdx == len(s.Data.Builds)-1 {
			s.UI.SelectedBuildIdx = 0
		} else {
			s.UI.SelectedBuildIdx++
		}
	}
	return nil
}

func reduceIdxPrev(s *AppState) []Effect {
	switch s.UI.FocusedView {
	case "projects":
		if len(s.Data.Projects) == 0 {
			return nil
		}
		if s.UI.SelectedProjectIdx == 0 {
			s.UI.SelectedProjectIdx = len(s.Data.Projects) - 1
		} else {
			s.UI.SelectedProjectIdx--
		}
	case "builds":
		if len(s.Data.Builds) == 0 {
			return nil
		}
		if s.UI.SelectedBuildIdx == 0 {
			s.UI.SelectedBuildIdx = len(s.Data.Builds) - 1
		} else {
			s.UI.SelectedBuildIdx--
		}
	}
	return nil
}

func reduceRenderHelp(s *AppState) []Effect {
	s.UI.ShowHelp = true
	return nil
}
