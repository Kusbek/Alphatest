package apiserver

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *server) handleUserDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		n, err := s.store.User().Delete(id)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		s.respond(w, r, http.StatusOK, map[string]int{"deleted": n})
	}
}
