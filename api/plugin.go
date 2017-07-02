package api

//go:generate protoc afanty.proto --go_out=plugins=grpc:.

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
