package domain

import (
	"demo/src/domain/entitie"
)
type UserRepository interface {
	Save(User *entitie.User) error
	GetAll() ([]entitie.User, error)
	Delete(id int) error
	Edit(id int, User *entitie.User) error
}