package apiserver

import (
	"alphatest/internal/store"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	errIncorrectEmailOrPassword = errors.New("Incorrect email or password")
	errNotAuthenticated         = errors.New("Not authenticated")
)

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {

	s.router.HandleFunc("/users", s.handleUserCreate()).Methods("POST")
	s.router.HandleFunc("/users/{id}", s.handleUserGet()).Methods("GET")

	admin := s.router.PathPrefix("/admin").Subrouter()
	admin.Use(s.authenticateAdmin)
	admin.HandleFunc("/users/{id}", s.handleUserDelete()).Methods("DELETE")
	admin.HandleFunc("/users", s.handleUsersGetByRole()).Methods("GET")

}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
