package application

import (
	"demo/src/domain"
	"demo/src/domain/entitie"
)
type ViewUsers struct {
	db domain.UserRepository
}

func NewViewUsers(db domain.UserRepository) *ViewUsers {
	return &ViewUsers{db: db}
}

func (v *ViewUsers) Execute() ([]entitie.User, error) {
	return v.db.GetAll()
}