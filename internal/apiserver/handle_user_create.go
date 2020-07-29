package apiserver

import (
	"alphatest/internal/model"
	"encoding/json"
	"net/http"
)

func (s *server) handleUserCreate() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Username: req.Username,
			Password: req.Password,
			RoleID:   1,
		}

		err := u.Validate()
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.RemovePassword()
		s.respond(w, r, http.StatusCreated, map[string]interface{}{"user": u})
	}
}
