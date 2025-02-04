package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type UserStorage struct {
	users []User
	mutex sync.Mutex
}


func NewUserStorage() *UserStorage {
	return &UserStorage{}
}

func (us *UserStorage) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &us.users)
}

func (us *UserStorage) CheckNewUsersHandler(c *gin.Context) {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	
	var newUsers []User

	for _, user := range us.users {
		if user.ID > 5 {
			newUsers = append(newUsers, user)
		}
	}

	if len(newUsers) > 0 {
		c.JSON(http.StatusOK, gin.H{"new_users": newUsers})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "No new users"})
	}
}
