package api

type (
	StateError struct {
		Name string
		Err  error
	}

	Run func(MInput) (MOutput, StateError)
)
