package state

type EffectKind string

const (
	DispatchLoadSession  EffectKind = "LoadSession"
	FetchProjects  EffectKind = "FetchProjects"
	CloseProgram  EffectKind = "CloseProgram"
)

type Effect struct {
  Kind    EffectKind
}
