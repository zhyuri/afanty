package server

import (
	"golang.org/x/net/context"
	"github.com/zhyuri/afanty/api"
)

type stateMachineServer struct {
}

func newStateMachineServer() *stateMachineServer {
	return new(stateMachineServer)
}

func (stateMachineServer) Run(ctx context.Context, input *api.MInput) (*api.MOutput, error) {

	return nil, nil
}
