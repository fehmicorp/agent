package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx    context.Context
	server *http.Server
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go a.startLocalServer()
}

func (a *App) startLocalServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/show", func(w http.ResponseWriter, r *http.Request) {
		runtime.WindowShow(a.ctx)
		runtime.WindowUnminimise(a.ctx)
		runtime.WindowCenter(a.ctx)
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/quit", func(w http.ResponseWriter, r *http.Request) {
		go runtime.Quit(a.ctx)
		w.WriteHeader(http.StatusOK)
	})
	a.server = &http.Server{
		Addr:    "127.0.0.1:8051",
		Handler: mux,
	}
	_ = a.server.ListenAndServe()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
