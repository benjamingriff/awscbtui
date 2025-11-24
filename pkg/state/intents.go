package state

type Intent interface {
	IntentName() string
}

type LoadSession struct{}
func (LoadSession) IntentName() string { return "LoadSession" }

type ViewNext struct{}
func (ViewNext) IntentName() string { return "ViewNext" }

type ViewPrev struct{}
func (ViewPrev) IntentName() string {
	return "ViewPrev"
}

type IdxNext struct{}
func (IdxNext) IntentName() string { return "IdxNext" }

type IdxPrev struct{}
func (IdxPrev) IntentName() string {
	return "IdxPrev"
}

type RenderHelp struct{}
func (RenderHelp) IntentName() string {
	return "RenderHelp"
}

type Quit struct{}
func (Quit) IntentName() string {
	return "Quit"
}
