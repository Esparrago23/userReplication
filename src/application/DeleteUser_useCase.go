package application

import (
	"demo/src/domain"
)
type DeleteUser struct {
	db domain.UserRepository
}

func NewDeleteUser(db domain.UserRepository) *DeleteUser {
	return &DeleteUser{db: db}
}

func (d *DeleteUser) Execute(id int) error {
	err:= d.db.Delete(id)
	if err != nil {
		return err
	}
	return nil
}