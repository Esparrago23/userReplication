package controllers

import (
	"demo/src/application"
	"demo/src/domain/entitie"
	"net/http"
	"github.com/gin-gonic/gin"
)
type CreateUserController struct {
	useCase application.CreateUser
}

func NewCreateUserController(useCase application.CreateUser) *CreateUserController {
	return &CreateUserController{useCase: useCase}
}

func (cu_c *CreateUserController) Execute(c *gin.Context) {
	var user struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		UserName string `json:"username"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := entitie.User{
		ID:       user.ID,
		Name:     user.Name,
		UserName: user.UserName,
	}

	err := cu_c.useCase.Execute(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}