package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func renderJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		serveError(w, err)
	}
}

func badRequest(w http.ResponseWriter) {
	http.Error(
		w,
		http.StatusText(http.StatusBadRequest),
		http.StatusBadRequest,
	)
}

func notFound(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func conflict(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
}

func serveError(w http.ResponseWriter, err error) {
	if err.Error() == "pg: no rows in result set" {
		notFound(w)
		return
	}

	if strings.Contains(err.Error(), "duplicate key value") {
		conflict(w)
		return
	}

	log.Printf("[ERROR] %s", err.Error())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
