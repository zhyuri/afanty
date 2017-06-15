package web

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"net/http"
	"context"
)

var srv *http.Server

func Run(addr string) {
	http.HandleFunc("/", handle)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	srv = &http.Server{
		Addr:    addr,
		Handler: nil,
	}
	logrus.Fatal(srv.ListenAndServe())
}

func Shutdown(ctx context.Context) {
	srv.Shutdown(ctx)
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
