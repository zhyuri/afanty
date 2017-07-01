package server

import (
	"github.com/zhyuri/afanty/api"
	"github.com/zhyuri/afanty/core"
	"golang.org/x/net/context"
)

type stateMachineServer struct {
	ACore *core.AfantyCore
}

func newStateMachineServer(c *core.AfantyCore) *stateMachineServer {
	return &stateMachineServer{
		ACore: c,
	}
}

func (stateMachineServer) Run(ctx context.Context, input *api.MInput) (*api.MOutput, error) {

	return nil, nil
}
