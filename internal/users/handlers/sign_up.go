package user

import (
	"go-task-app/internal/users/helpers"
	"go-task-app/internal/users/services"
	"go-task-app/internal/users/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var signUpInput types.SignUpInput

	err := c.ShouldBindJSON(&signUpInput)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "入力値が不正です。",
		})
		return
	}

	result, err := services.SignUp(&signUpInput)

	if err != nil {
		helpers.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": result,
	})
}
