package server

import (
	"fmt"
	"net/http"
	"time"
)

type App struct {
	httpServer *http.Server
}

func (a *App) Run(port string) error {
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	http.HandleFunc("/", HelloServer)

	return a.httpServer.ListenAndServe()
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World %s!", r.URL.Path[1:])
}
