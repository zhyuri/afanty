package server

import (
	"github.com/zhyuri/afanty/api"
	"golang.org/x/net/context"
)

type stateMachineServer struct {
}

func newStateMachineServer() *stateMachineServer {
	return new(stateMachineServer)
}

func (stateMachineServer) Run(ctx context.Context, input *api.MInput) (*api.MOutput, error) {

	return nil, nil
}
