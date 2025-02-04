package infrastructure

import (
	"demo/src/domain/entitie"
	"fmt"
	"log"
	"sync"
)

type UserStorage struct {
	mu    sync.Mutex
	users []entitie.User
}

func NewUserStorage() *UserStorage {
	return &UserStorage{
		users: []entitie.User{},
	}
}

func (us *UserStorage) Save(user *entitie.User) error {
	us.mu.Lock()
	defer us.mu.Unlock()

	if len(us.users) == 0 {
		user.ID = 1
	} else {
		user.ID = us.users[len(us.users)-1].ID + 1
	}

	us.users = append(us.users, *user)
	fmt.Println("[Storage] - Usuario guardado:", user)
	return nil
}

func (us *UserStorage) GetAll() ([]entitie.User, error) {
	us.mu.Lock()
	defer us.mu.Unlock()

	fmt.Println("[Storage] - Lista de usuarios")
	return us.users, nil
}

func (us *UserStorage) Edit(id int, updatedUser *entitie.User) error {
	us.mu.Lock()
	defer us.mu.Unlock()

	for i, user := range us.users {
		if user.ID == id {
			us.users[i].Name = updatedUser.Name
			us.users[i].UserName = updatedUser.UserName
			fmt.Println("[Storage] - Usuario actualizado:", updatedUser)
			return nil
		}
	}
	log.Printf("[Storage] - Usuario con ID %d no encontrado", id)
	return fmt.Errorf("usuario no encontrado")
}

func (us *UserStorage) Delete(id int) error {
	us.mu.Lock()
	defer us.mu.Unlock()

	for i, user := range us.users {
		if user.ID == id {
			us.users = append(us.users[:i], us.users[i+1:]...)
			fmt.Println("[Storage] - Usuario eliminado con ID:", id)
			return nil
		}
	}
	log.Printf("[Storage] - Usuario con ID %d no encontrado", id)
	return fmt.Errorf("usuario no encontrado")
}
