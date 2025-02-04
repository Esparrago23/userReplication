package controllers

import (
	"demo/src/application"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
)
type DeleteUserController struct {
	useCase application.DeleteUser
}

func NewDeleteUserController(useCase application.DeleteUser) *DeleteUserController {
	return &DeleteUserController{useCase: useCase}
}

func (du_c *DeleteUserController) Execute(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.Atoi(productID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }


	err = du_c.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}