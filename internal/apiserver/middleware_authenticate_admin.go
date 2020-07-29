package apiserver

import (
	"net/http"
)

func (s *server) authenticateAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			adminHeader := r.Header.Get("Authorization")
			if adminHeader != "IMADMIN" {
				s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
				return
			}
			next.ServeHTTP(w, r)
		})
}
