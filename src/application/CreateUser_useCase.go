package application

import (
	"demo/src/domain"
	"demo/src/domain/entitie"
)
type CreateUser struct{
	db domain.UserRepository
}
func NewCreateUser(db domain.UserRepository) *CreateUser {
	return &CreateUser{db: db}
}
func (c *CreateUser) Execute(NewUser *entitie.User) error {
	err := c.db.Save(NewUser)
	if err != nil {
		return err
	}
	return nil
}