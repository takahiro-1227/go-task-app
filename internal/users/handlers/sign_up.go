package user

import (
	"github.com/gin-gonic/gin"
	"go-task-app/internal/users/services"
	"go-task-app/internal/users/types"
	"net/http"
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

	result, code, err := services.SignUp(&newUser)

	if err != nil {
		c.JSON(code, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(code, gin.H{
		"user": result,
	})
}
