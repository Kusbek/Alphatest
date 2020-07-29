package apiserver

import (
	"alphatest/internal/model"
	"alphatest/internal/store/sqlstore"
	"database/sql"
	"net/http"
)

//Start ...
func Start(config *Config) error {
	db, err := NewDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(1)
	defer db.Close()
	store := sqlstore.New(db)
	srv := newServer(store)
	srv.store.User().Create(&model.User{
		Username: "admin",
		Password: "admin",
		RoleID:   2,
	})
	return http.ListenAndServe(config.BindAddr, srv)
}

//NewDB ...
func NewDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
