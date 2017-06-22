package web

import (
	"context"
	"fmt"
	"github.com/Sirupsen/logrus"
	"net/http"
)

var srv *http.Server

func Run(addr string) {
	http.HandleFunc("/", handle)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	srv = &http.Server{
		Addr:    addr,
		Handler: nil,
	}
	logrus.Infoln("HTTP Listening on port ", addr)
	logrus.Fatal(srv.ListenAndServe())
}

func Shutdown(ctx context.Context) {
	srv.Shutdown(ctx)
	logrus.Infoln("HTTP shut down")
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello Afanty!")
}

func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "ok")
}
