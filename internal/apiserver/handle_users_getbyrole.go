package apiserver

import (
	"alphatest/internal/model"
	"errors"
	"net/http"
	"strconv"
)

func (s *server) handleUsersGetByRole() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roleIDStr := r.URL.Query().Get("role_id")
		if roleIDStr == "" {
			roleIDStr = "1"
		}
		roleIDInt, err := strconv.Atoi(roleIDStr)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errors.New("role_id is invalid"))
			return
		}

		pageNumStr := r.URL.Query().Get("page_num")
		if pageNumStr == "" {
			pageNumStr = "1"
		}
		pageNumInt, err := strconv.Atoi(pageNumStr)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errors.New("page_num is invalid"))
			return
		}

		pageSizeStr := r.URL.Query().Get("page_size")
		if pageSizeStr == "" {
			pageSizeStr = "5"
		}

		pageSizeInt, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errors.New("page_size is invalid"))
			return
		}

		users, count, err := s.store.User().GetByRole(roleIDInt, pageNumInt, pageSizeInt)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		links := model.NewLinks(r, count, pageNumInt, pageSizeInt)
		s.respond(w, r, http.StatusOK, map[string]interface{}{
			"users": users,
			"links": links,
		})
	}
}
