package apiserver

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *server) handleUserGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().Find(id)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		s.respond(w, r, http.StatusOK, map[string]interface{}{"user": u})
	}
}
