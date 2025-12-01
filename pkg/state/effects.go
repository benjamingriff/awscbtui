package state

type EffectKind string

const (
	DispatchLoadSession  EffectKind = "LoadSession"
	FetchProjects  EffectKind = "FetchProjects"
	LoadProjectsBuilds  EffectKind = "LoadProjectsBuilds"
	CloseProgram  EffectKind = "CloseProgram"
)

type Effect struct {
  Kind    EffectKind
	Data 		any
}

type LoadProjectsBuildsData struct {
  ProjectName string
}
