package main

import (
	"demo/users"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userStorage := users.NewUserStorage()

	err := userStorage.LoadFromFile("users.json")
	if err != nil {
		log.Fatalf("Error cargando usuarios: %v", err)
	}

	router.GET("/check-new-users", userStorage.CheckNewUsersHandler)

	router.Run(":8080")
}
