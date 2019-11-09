package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/justcompile/midgard/common/dal"

	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"
)

type projects struct {
	db *pg.DB
}

func (p *projects) List(w http.ResponseWriter, r *http.Request) {
	results := make([]*dal.Project, 0)

	err := p.db.Model(&results).
		OrderExpr("created_at DESC").
		Select()

	if err != nil {
		serveError(w, err)
		return
	}

	renderJSON(w, results)
}

func (p *projects) Get(w http.ResponseWriter, r *http.Request) {
	idFromURL := chi.URLParam(r, "id")

	if idFromURL == "" {
		badRequest(w)
		return
	}

	var err error
	var id int

	if id, err = strconv.Atoi(idFromURL); err != nil {
		badRequest(w)
		return
	}

	project := &dal.Project{Id: int64(id)}
	err = p.db.Select(project)
	if err != nil {
		serveError(w, err)
		return
	}

	renderJSON(w, project)
}

func (p *projects) Create(w http.ResponseWriter, r *http.Request) {
	var project *dal.Project

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		log.Printf("[ERROR] %s", err.Error())
		badRequest(w)
		return
	}

	if _, err := p.db.Model(project).Returning("*").Insert(); err != nil {
		serveError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	renderJSON(w, project)
}

func (p *projects) Update(w http.ResponseWriter, r *http.Request) {
	idFromURL := chi.URLParam(r, "id")

	if idFromURL == "" {
		badRequest(w)
		return
	}

	var err error
	var id int

	if id, err = strconv.Atoi(idFromURL); err != nil {
		badRequest(w)
		return
	}

	project := &dal.Project{Id: int64(id)}
	err = p.db.Select(project)
	if err != nil {
		serveError(w, err)
		return
	}

	var projToUpdate *dal.Project

	if err := json.NewDecoder(r.Body).Decode(&projToUpdate); err != nil {
		log.Printf("[ERROR] %s", err.Error())
		badRequest(w)
		return
	}

	projToUpdate.Id = project.Id
	projToUpdate.CreatedAt = project.CreatedAt

	if err := p.db.Update(projToUpdate); err != nil {
		serveError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	renderJSON(w, projToUpdate)
}

func (p *projects) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", p.List)
	r.Post("/", p.Create)
	r.Get("/{id}", p.Get)
	r.Put("/{id}", p.Update)

	return r
}
