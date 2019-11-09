package internal

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/justcompile/midgard/web/api"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/api", api.Routes())

	return r
}
