package model

import (
	"fmt"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

//User ...
type User struct {
	ID                int    `json:"id,omitempty"`
	Username          string `json:"username,omitempty"`
	Password          string `json:"-"`
	EncryptedPassword string `json:"-"`
	Role              string `json:"role,omitempty"`
	RoleID            int    `json:"-"`
}

//Validate ...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.Password, validation.Length(6, 100)),
	)
}

//RemovePassword ...
func (u *User) RemovePassword() {
	u.Password = ""
}

//EncryptPassword ...
func (u *User) EncryptPassword() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}
	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

//TestUsers ...
func TestUsers(t *testing.T) []*User {
	users := make([]*User, 0)

	users = append(users, &User{
		ID:       1,
		Username: "kusbek123",
		Password: "password",
		RoleID:   2,
	})

	for i := 2; i < 100; i++ {
		users = append(users, &User{
			ID:       i,
			Username: fmt.Sprintf("testuser%d", i),
			Password: fmt.Sprintf("password%d", i),
			RoleID:   1,
		})
	}
	return users
}
