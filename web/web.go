package web

import (
	"context"
	"fmt"
	"github.com/Sirupsen/logrus"
	"net/http"
)

var srv *http.Server

func Run() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	logrus.Print("Listening on port 8080")
	srv = &http.Server{
		Addr:    ":8080",
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
