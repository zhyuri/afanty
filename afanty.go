package main

import (
	"context"
	"github.com/Sirupsen/logrus"
	"github.com/zhyuri/afanty/web"
	"os"
	"os/signal"
	"time"
)

func main() {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	go web.Run()

	<-stopChan
	logrus.Infoln("Shutting down server...")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	web.Shutdown(ctx)
	logrus.Infoln("Server gracefully stopped")

}
