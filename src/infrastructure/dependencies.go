package infrastructure

import (
	"demo/src/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	Create *controllers.CreateUserController
	Get    *controllers.GetUsersController
	Edit   *controllers.EditUserController
	Delete *controllers.DeleteUserController
}

func UsersRoutes(router *gin.Engine, handlers UserHandlers) {
	usersGroup := router.Group("/users")
	{
		usersGroup.POST("/", handlers.Create.Execute)
		usersGroup.GET("/", handlers.Get.Execute)
		usersGroup.PUT("/:id", handlers.Edit.Execute)
		usersGroup.DELETE("/:id", handlers.Delete.Execute)
	}
}