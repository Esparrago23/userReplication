package infrastructure

import (
	"demo/src/application"
	"demo/src/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	us := NewUserStorage()

	createUserService := application.NewCreateUser(us)
	viewUserService := application.NewViewUsers(us)
	editUserService := application.NewEditUser(us)
	deleteUserService := application.NewDeleteUser(us)

	createUserController := controllers.NewCreateUserController(*createUserService)
	viewUserController := controllers.NewGetUsersController(*viewUserService)
	editUserController := controllers.NewEditUserController(*editUserService)
	deleteUserController := controllers.NewDeleteUserController(*deleteUserService)

	UsersRoutes(router, UserHandlers{
		Create: createUserController,
		Get:    viewUserController,
		Edit:   editUserController,
		Delete: deleteUserController,
	})
}
