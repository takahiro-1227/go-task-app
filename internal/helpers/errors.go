package helpers

import (
	"log"
	"net/http"

	"go-task-app/internal/constants"

	"github.com/gin-gonic/gin"
)

func HandleHandlerError(c *gin.Context, err error) {
	log.Println(err.Error())

	switch err {
	case constants.ErrInvalidInput:
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "予期せぬエラーが発生しました。",
		})
	}
}
