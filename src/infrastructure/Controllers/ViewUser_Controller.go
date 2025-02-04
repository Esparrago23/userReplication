package controllers

import (
	"demo/src/application"
	
	"net/http"
	"github.com/gin-gonic/gin"
)
type GetUsersController struct {
	useCase application.ViewUsers
}

func NewGetUsersController(useCase application.ViewUsers) *GetUsersController {
	return &GetUsersController{useCase: useCase}
}

func (gu_c *GetUsersController) Execute(c *gin.Context) {
	users, err := gu_c.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}