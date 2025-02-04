package controllers

import (
	"demo/src/application"
	"demo/src/domain/entitie"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)
type EditUserController struct {
	useCase application.EditUser
}

func NewEditUserController(useCase application.EditUser) *EditUserController {
	return &EditUserController{useCase: useCase}
}

func (eu_c *EditUserController) Execute(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user struct {
		Name     string `json:"name"`
		UserName string `json:"username"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser := entitie.User{
		ID:       id,
		Name:     user.Name,
		UserName: user.UserName,
	}

	err = eu_c.useCase.Execute(id, &updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}