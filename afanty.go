package main

import (
	"context"
	"github.com/Sirupsen/logrus"
	"github.com/zhyuri/afanty/web"
	"os"
	"os/signal"
	"time"
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

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = ":8080"
	}

	go web.Run(port)
	logrus.Infoln("Listening on port ", port)

	<-stopChan
	logrus.Infoln("Shutting down server...")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	web.Shutdown(ctx)
	logrus.Infoln("Server gracefully stopped")
}
