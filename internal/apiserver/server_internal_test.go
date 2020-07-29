package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"alphatest/internal/apiserver/teststore"
	"alphatest/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"username": "kusbek",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid  payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "invalid",
			},
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, tc.expectedCode)
		})
	}
}

func TestServer_HandleUserGet(t *testing.T) {
	s := newServer(teststore.New())
	users := model.TestUsers(t)

	for _, user := range users {
		s.store.User().Create(user)
		t.Run(user.Username, func(t *testing.T) {
			rec := httptest.NewRecorder()
			url := fmt.Sprintf("/users/%d", len(user.Username))
			req, _ := http.NewRequest(http.MethodGet, url, nil)
			s.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}

	rec := httptest.NewRecorder()
	url := fmt.Sprintf("/users/%d", 5)
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	s.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)

	rec = httptest.NewRecorder()
	url = fmt.Sprintf("/users/%s", "invalid")
	req, _ = http.NewRequest(http.MethodGet, url, nil)
	s.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestServer_HandleUsersGetByRole(t *testing.T) {
	s := newServer(teststore.New())
	users := model.TestUsers(t)

	for _, user := range users {
		s.store.User().Create(user)
	}

	testCases := []struct {
		name          string
		query         string
		expectedCode  int
		expectedCount int
	}{
		{
			name:         "valid 2",
			query:        "?role_id=2",
			expectedCode: http.StatusOK,
		},
		{
			name:         "valid 1",
			query:        "?role_id=1",
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalivalid",
			query:        "?role_id=A",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/admin/users"+tc.query, nil)
			req.Header.Add("Authorization", "IMADMIN")
			s.ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, tc.expectedCode)
		})
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/admin/users", nil)
	s.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/admin/users", nil)
	req.Header.Add("Authorization", "Invalid")
	s.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestServer_HandleUserDelete(t *testing.T) {
	s := newServer(teststore.New())
	users := model.TestUsers(t)

	for _, user := range users {
		s.store.User().Create(user)
		t.Run(user.Username, func(t *testing.T) {
			rec := httptest.NewRecorder()
			url := fmt.Sprintf("/admin/users/%d", len(user.Username))
			req, _ := http.NewRequest(http.MethodDelete, url, nil)
			req.Header.Add("Authorization", "IMADMIN")
			s.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
		})
	}
}
