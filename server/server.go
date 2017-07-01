package server

import (
	"context"
	"net"

	"github.com/Sirupsen/logrus"
	"github.com/zhyuri/afanty/api"
	"github.com/zhyuri/afanty/core"
	"google.golang.org/grpc"
)

var grpcServer *grpc.Server

func Run(c *core.AfantyCore, addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer = grpc.NewServer(opts...)

	api.RegisterStateMachineServer(grpcServer, newStateMachineServer(c))

	logrus.Infoln("RPC Listening on port ", addr)
	grpcServer.Serve(lis)
}

func Shutdown(ctx context.Context) {
	grpcServer.GracefulStop()
	logrus.Infoln("RPC shut down")
}
