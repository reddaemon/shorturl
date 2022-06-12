package router

import (
	"shorturl/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Route("/v1/shorturl", func(r chi.Router) {
		r.Get("/short", handlers.ShortHandler)
	})

	return r
}
