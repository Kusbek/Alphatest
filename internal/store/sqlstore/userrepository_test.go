package sqlstore_test

import (
	"alphatest/internal/model"
	"alphatest/internal/store/sqlstore"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	if databaseURL == "" {
		databaseURL = "host=localhost user=postgres password=1234 dbname=restapi_dev sslmode=disable"
	}
	os.Exit(m.Run())
}

func clear(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		return err
	}
	return nil
}

//TestDB  ...
func TestUserRepository_Create(t *testing.T) {
	db := sqlstore.TestDB(t, databaseURL)
	s := sqlstore.New(db)
	users := model.TestUsers(t)

	for _, user := range users {
		err := s.User().Create(user)
		assert.NoError(t, err)
		assert.NotNil(t, user)
	}
	err := clear(db)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.Close()
}

func TestUserRepository_Find(t *testing.T) {
	db := sqlstore.TestDB(t, databaseURL)

	s := sqlstore.New(db)
	users := model.TestUsers(t)
	for _, user := range users {
		err := s.User().Create(user)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	for _, user := range users {
		u, err := s.User().Find(user.ID)
		assert.NoError(t, err)
		assert.NotNil(t, u)
	}
	err := clear(db)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.Close()
}

func TestUserRepository_FindByRole(t *testing.T) {
	db := sqlstore.TestDB(t, databaseURL)

	s := sqlstore.New(db)
	users := model.TestUsers(t)
	for _, user := range users {
		err := s.User().Create(user)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	admin, _, err := s.User().GetByRole(2, 1, 5)
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	assert.Equal(t, 1, len(admin))

	pageSize := 5
	res, count, err := s.User().GetByRole(1, 1, pageSize)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(len(res))
	assert.NoError(t, err)
	assert.Equal(t, len(users[1:]), count)
	assert.Equal(t, pageSize, len(res))

	err = clear(db)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.Close()
}

func TestUserRepository_Delete(t *testing.T) {
	db := sqlstore.TestDB(t, databaseURL)

	s := sqlstore.New(db)
	users := model.TestUsers(t)
	for _, user := range users {
		err := s.User().Create(user)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	for _, user := range users {
		n, err := s.User().Delete(user.ID)
		if err != nil {
			log.Fatal(err.Error())
		}

		assert.Equal(t, 1, n)
	}

	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM users")
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, 1, count)
}
