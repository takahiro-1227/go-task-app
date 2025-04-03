package helpers

import (
	"go-task-app/internal/users/constants"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	log.Println(err.Error())

	switch err {
	case constants.ErrSignIn:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	case constants.ErrSignInServer:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	default:
		if strings.Contains(err.Error(), "は必須です。") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "予期せぬエラーが発生しました。",
		})
	}
}
