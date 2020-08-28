package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	gopher "github.com/srehouni/api-go/pkg"
)

type api struct {
	router     http.Handler
	repository gopher.GopherRepository
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()

	r.HandleFunc("/", a.fetchGophers).Methods(http.MethodGet)
	r.HandleFunc("/gophers", a.fetchGophers).Methods(http.MethodGet)
	r.HandleFunc("/gophers/{ID:[a-zA-Z0-9_]+}", a.fetchGopher).Methods(http.MethodGet)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) fetchGophers(w http.ResponseWriter, r *http.Request) {

	gophers, _ := a.repository.FetchGophers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gophers)
}

func (a *api) fetchGopher(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	gopher, err := a.repository.FetchGopherByID(vars["ID"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Gopher Not found")
		return
	}

	json.NewEncoder(w).Encode(gopher)
}
