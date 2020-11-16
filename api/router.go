package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// BuildRouter builds the router
func BuildRouter(app *App) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Heartbeat("/ping"))

	// Routes

	// POST /save
	mux.Post("/save", app.HandleSaveText)

	return mux
}
