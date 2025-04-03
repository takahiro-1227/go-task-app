package user

import (
	"go-task-app/internal/users/helpers"
	"go-task-app/internal/users/services"
	"go-task-app/internal/users/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var signInInput types.SignInInput

	err := c.ShouldBindJSON(&signInInput)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "入力値が不正です。",
		})
		return
	}

	result, err := services.SignIn(signInInput)

	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	c.JSON(200, result)
}
