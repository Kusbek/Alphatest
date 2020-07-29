package teststore

import (
	"alphatest/internal/model"
	"alphatest/internal/store"
)

//UserRepository ...
type UserRepository struct {
	store *Store
	users map[int]*model.User
}

//Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.EncryptPassword(); err != nil {
		return err
	}
	u.ID = len(u.Username)
	r.users[u.ID] = u

	return nil
}

//Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

//GetByRole ...
func (r *UserRepository) GetByRole(roleID, pageNum, pageSize int) ([]*model.User, int, error) {
	result := make([]*model.User, 0)
	for _, user := range r.users {
		if user.RoleID == roleID {
			result = append(result, user)
		}
	}

	return result, len(result), nil
}

//Delete ...
func (r *UserRepository) Delete(id int) (int, error) {

	_, ok := r.users[id]
	if !ok {
		return 0, nil
	}
	delete(r.users, id)
	return 1, nil
}
