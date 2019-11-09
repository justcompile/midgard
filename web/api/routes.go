package api

import (
	"github.com/justcompile/midgard/common/dal"

	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	db := dal.Database()

	projects := &projects{db}
	workers := &workers{}

	r.Mount("/projects", projects.routes())
	r.Mount("/workers", workers.routes())

	return r
}
