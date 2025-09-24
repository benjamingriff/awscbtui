package state

type Intent interface {
	Name() string
}

type FocusNext struct{}

func (FocusNext) Name() string {
	return "FocusNext"
}

type FocusPrev struct{}

func (FocusPrev) Name() string {
	return "FocusPrev"
}

type MoveFocus struct {
    View View
    Delta int
}
