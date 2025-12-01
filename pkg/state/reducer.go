package state

func ReduceIntent(s *AppState, i Intent) []Effect {
	switch i.(type) {
	case LoadSession:
		return reduceLoadSession(s)
	case ViewNext:
		return reduceViewNext(s)
	case ViewPrev:
		return reduceViewPrev(s)
	case IdxNext:
		return reduceIdxNext(s)
	case IdxPrev:
		return reduceIdxPrev(s)
	case RenderHelp:
		return reduceRenderHelp(s)
	case Quit:
		return reduceQuit(s)
	default:
		return nil
	}
}

func ReduceMessage(s *AppState, m Message) []Effect {
	switch v := m.(type) {
	case SessionLoaded:
		return reduceSessionLoaded(s, v)
	case ProjectsLoaded:
		return reduceProjectsLoaded(s, v)
	case BuildsLoaded:
		return reduceBuildsLoaded(s, v)
	default:
		return nil
	}
}
