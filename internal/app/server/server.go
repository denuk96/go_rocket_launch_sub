package server

import (
	"context"
	"net/http"
	"time"
)

type App struct {
	httpServer *http.Server
}

func (a *App) Run(port string, handler http.Handler) error {
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return a.httpServer.ListenAndServe()
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.httpServer.Shutdown(ctx)
}
