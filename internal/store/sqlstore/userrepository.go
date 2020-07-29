package sqlstore

import (
	"alphatest/internal/model"
	"alphatest/internal/store"
	"database/sql"
)

//UserRepository ...
type UserRepository struct {
	store *Store
}

//Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.EncryptPassword(); err != nil {
		return err
	}
	err := r.store.db.QueryRow(
		"INSERT INTO users (username, encrypted_password, role_id) VALUES ($1, $2, $3) RETURNING id",
		u.Username,
		u.EncryptedPassword,
		u.RoleID,
	).Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}

//Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, username, encrypted_password FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Username,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

//GetByRole ...
func (r *UserRepository) GetByRole(roleID, pageNum, pageSize int) ([]*model.User, int, error) {
	users := make([]*model.User, 0)
	rows, err := r.store.db.Query(
		`
		SELECT users.id, users.username, roles.role_name, count(*) OVER()
		FROM users
		INNER JOIN roles ON users.role_id=roles.role_id
		WHERE users.role_id = $1 LIMIT $2 OFFSET $3;
	`, roleID, pageSize, (pageNum-1)*pageSize)
	if err != nil {
		return nil, 0, err
	}
	var count int
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Role,
			&count,
		)
		if err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}

	return users, count, nil
}

//Delete ...
func (r *UserRepository) Delete(id int) (int, error) {
	res, err := r.store.db.Exec("DELETE FROM users WHERE id=$1 AND role_id=1", id)
	if err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}
