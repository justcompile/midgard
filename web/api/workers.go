package api

import (
	"net/http"

	"github.com/justcompile/midgard/common"

	"github.com/go-chi/chi"
)

type workers struct{}

func (w *workers) List(resp http.ResponseWriter, r *http.Request) {
	renderJSON(resp, common.WorkerRegistry.GetConnectedWorkers())
}

func (w *workers) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", w.List)

	return r
}
