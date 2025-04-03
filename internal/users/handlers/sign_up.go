package user

import (
	"go-task-app/internal/users/helpers"
	"go-task-app/internal/users/services"
	"go-task-app/internal/users/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var newUser types.User

	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "入力値が不正です。",
		})
		return
	}

	result, err := services.SignUp(&newUser)

	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"user": result,
	})
}
