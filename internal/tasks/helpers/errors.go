package helpers

import (
	"go-task-app/internal/tasks/constants"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleTaskError(c *gin.Context, err error) {
	log.Println(err)

	switch err {
	case constants.ErrTitleIsEmpty, constants.ErrInvalidUpdate, constants.ErrInvalidDelete:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	case constants.ErrGetTasks, constants.ErrCreateFailed, constants.ErrUpdateFailed, constants.ErrDeleteFailed:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "予期せぬエラーが発生しました。",
		})
	}
}
