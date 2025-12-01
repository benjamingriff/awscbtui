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
	return []Effect{{Kind: LoadProjectsBuildIds, Data: LoadProjectsBuildIdsData{ProjectName: firstProject.Name}}}
}

func reduceBuildIdsLoaded(s *AppState, m BuildIdsLoaded) []Effect {
	s.Data.BuildIds[m.ProjectName] = m.BuildIds
return []Effect{{Kind: LoadProjectsBuilds, Data: LoadProjectsBuildsData{ProjectName: m.ProjectName, BuildIds: m.BuildIds}}}
}

func reduceBuildsLoaded(s *AppState, m BuildsLoaded) []Effect {
	s.Data.Builds[m.ProjectName] = m.Builds
	return nil
}
