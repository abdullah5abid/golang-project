package dao

import "github.com/spongeling/admin-api/internal/errors"

type User struct {
	Id       uint64 `db:"id"`
	Username string `db:"name"`
	Password string `db:"password"`
}

func (u *User) GetID() uint64 {
	return u.Id
}

func (*User) GetTable() string {
	return "user"
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.BadRequest("username is required")
	}
	if u.Password == "" {
		return errors.BadRequest("password is required")
	}
	return nil
}
