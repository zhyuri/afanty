package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/zhyuri/afanty/core"
	"github.com/zhyuri/afanty/server"
	"github.com/zhyuri/afanty/web"
)

var (
	gitTag    string
	buildTime string
)

func main() {
	logrus.Infoln("gitTig: ", gitTag)
	logrus.Infoln("BuildTime: ", buildTime)

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = ":8080"
	}
	rpcPort := os.Getenv("RPC_PORT")
	if rpcPort == "" {
		rpcPort = ":10043"
	}

	c := core.NewAfantyInstance("main")

	go server.Run(c, rpcPort)
	go web.Run(httpPort)

	<-stopChan
	logrus.Infoln("Shutting down server...")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	web.Shutdown(ctx)
	server.Shutdown(ctx)

	<-ctx.Done()
	if err := ctx.Err(); err != nil {
		logrus.Infoln("Server stopped error", err)
	} else {
		logrus.Infoln("Server gracefully stopped")
	}
}
