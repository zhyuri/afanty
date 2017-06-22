package server

import (
	"context"
	"github.com/Sirupsen/logrus"
	"github.com/zhyuri/afanty/api"
	"google.golang.org/grpc"
	"net"
)

var grpcServer *grpc.Server

func Run(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer = grpc.NewServer(opts...)

	api.RegisterStateMachineServer(grpcServer, newStateMachineServer())

	logrus.Infoln("RPC Listening on port ", addr)
	grpcServer.Serve(lis)
}

func Shutdown(ctx context.Context) {
	grpcServer.GracefulStop()
	logrus.Infoln("RPC shut down")
}
