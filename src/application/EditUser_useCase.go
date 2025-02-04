package application

import (
	"demo/src/domain"
	"demo/src/domain/entitie"
)

type EditUser struct {
	db domain.UserRepository
}

func NewEditUser(db domain.UserRepository) *EditUser {
	return &EditUser{db: db}
}

func (e *EditUser) Execute(id int, user *entitie.User) error {
	return e.db.Edit(id, user)
}