package api

type (
	StateError struct {
		Name string
		Err  error
	}

	Run func(MInput) (MOutput, StateError)
)

func (e StateError) Error() string {
	return e.Name + ":" + e.Err.Error()
}
