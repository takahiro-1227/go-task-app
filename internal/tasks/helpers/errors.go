package helpers

import (
	"go-task-app/internal/tasks/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleTaskError(c *gin.Context, err error) {
	switch err {
	case constants.ErrTitleIsEmpty, constants.ErrDuplicatedTitle:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	case constants.ErrCreateFailed:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "予期せぬエラーが発生しました。",
		})
	}
}
