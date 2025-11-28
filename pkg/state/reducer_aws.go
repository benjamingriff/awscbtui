package state

func reduceLoadSession(s *AppState) []Effect {
	s.UI.Loading = true
	return []Effect{{Kind: DispatchLoadSession}}
}

func reduceSessionLoaded(s *AppState, m SessionLoaded) []Effect {
	s.Session.Profile = m.SessionInfo.Profile
	return nil
}

func reduceProjectsLoaded(s *AppState, m ProjectsLoaded) []Effect {
	s.Data.Projects = m.Projects
	return nil
}
