package runtime

type State struct {
	HadError bool
}

func NewState() State {
	return State{
		HadError: false,
	}
}
