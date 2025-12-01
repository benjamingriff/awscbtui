package state

func reduceLoadSession(s *AppState) []Effect {
	s.UI.Loading = true
	return []Effect{{Kind: DispatchLoadSession}}
}

func reduceSessionLoaded(s *AppState, m SessionLoaded) []Effect {
	s.Session.Profile = m.SessionInfo.Profile
	s.UI.Loading = true
return []Effect{{Kind: FetchProjects}}
}

func reduceProjectsLoaded(s *AppState, m ProjectsLoaded) []Effect {
	s.Data.Projects = m.Projects
	firstProject := m.Projects[0]
	return []Effect{{Kind: LoadProjectsBuilds, Data: LoadProjectsBuildsData{ProjectName: firstProject.Name}}}
}


func reduceBuildsLoaded(s *AppState, m BuildsLoaded) []Effect {
	s.Data.Builds[m.ProjectName] = m.Builds
	return nil
}
