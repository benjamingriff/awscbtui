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
		if s.UI.FocusedProjectIdx == len(s.Data.Projects)-1 {
			s.UI.FocusedProjectIdx = 0
		} else {
			s.UI.FocusedProjectIdx++
		}
	case "builds":
		if len(s.Data.Builds[s.UI.SelectedProjectName]) == 0 {
			return nil
		}
		if s.UI.FocusedBuildIdx == len(s.Data.Builds[s.UI.SelectedProjectName])-1 {
			s.UI.FocusedBuildIdx = 0
		} else {
			s.UI.FocusedBuildIdx++
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
		if s.UI.FocusedProjectIdx == 0 {
			s.UI.FocusedProjectIdx = len(s.Data.Projects) - 1
		} else {
			s.UI.FocusedProjectIdx--
		}
	case "builds":
		if len(s.Data.Builds[s.UI.SelectedProjectName]) == 0 {
			return nil
		}
		if s.UI.FocusedBuildIdx == 0 {
			s.UI.FocusedBuildIdx = len(s.Data.Builds[s.UI.SelectedProjectName]) - 1
		} else {
			s.UI.FocusedBuildIdx--
		}
	}
	return nil
}

func reduceRenderHelp(s *AppState) []Effect {
	s.UI.ShowHelp = true
	return nil
}

func reduceMakeSelection(s *AppState) []Effect {
	switch s.UI.FocusedView {
	case "projects":
		s.UI.SelectedProjectIdx = s.UI.FocusedProjectIdx
		s.UI.SelectedProjectName = s.Data.Projects[s.UI.FocusedProjectIdx].Name
	case "builds":
		s.UI.SelectedBuildIdx = s.UI.FocusedBuildIdx
	}
	return nil
}
