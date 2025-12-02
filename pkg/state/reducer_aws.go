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
	s.UI.FocusedProjectIdx = 0
	s.UI.SelectedProjectName = firstProject.Name
	s.UI.SelectedProjectIdx = 0
	s.UI.Loading = true

	var effects []Effect
	for _, p := range m.Projects {
		effects = append(effects, Effect{Kind: LoadProjectsBuildIds, Data: LoadProjectsBuildIdsData{ProjectName: p.Name}})
	}
	return effects
}

func reduceBuildIdsLoaded(s *AppState, m BuildIdsLoaded) []Effect {
	s.Data.BuildIds[m.ProjectName] = m.BuildIds
	s.UI.Loading = true
return []Effect{{Kind: LoadProjectsBuilds, Data: LoadProjectsBuildsData{ProjectName: m.ProjectName, BuildIds: m.BuildIds}}}
}

func reduceBuildsLoaded(s *AppState, m BuildsLoaded) []Effect {
	s.Data.Builds[m.ProjectName] = m.Builds
	s.UI.Loading = false

	if s.UI.SelectedProjectName == m.ProjectName &&
	s.UI.FocusedBuildIdx == -1 &&
	len(m.Builds) > 0 {
		s.UI.FocusedBuildIdx = 0
		s.UI.SelectedBuildIdx = 0
		s.UI.SelectedBuildID = m.Builds[0].ID
	}
	return nil
}
