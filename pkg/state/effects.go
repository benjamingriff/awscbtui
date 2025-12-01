package state

type EffectKind string

const (
	DispatchLoadSession  EffectKind = "LoadSession"
	FetchProjects  EffectKind = "FetchProjects"
	LoadProjectsBuildIds  EffectKind = "LoadProjectsBuildIds"
	LoadProjectsBuilds  EffectKind = "LoadProjectsBuilds"
	CloseProgram  EffectKind = "CloseProgram"
)

type Effect struct {
  Kind    EffectKind
	Data 		any
}

type LoadProjectsBuildIdsData struct {
  ProjectName string
}

type LoadProjectsBuildsData struct {
	ProjectName string
  BuildIds []BuildId
}
