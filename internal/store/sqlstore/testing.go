package sqlstore

import (
	"database/sql"
	"testing"
)

//TestDB ...
func TestDB(t *testing.T, databaseURL string) *sql.DB {
	t.Helper()
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db
}
