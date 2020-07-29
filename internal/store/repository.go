package store

import "alphatest/internal/model"

//UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	GetByRole(int, int, int) ([]*model.User, int, error)
	Find(int) (*model.User, error)
	Delete(int) (int, error)
}
