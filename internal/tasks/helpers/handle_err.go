package helpers

import (
	"go-task-app/internal/tasks/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleTaskError(c *gin.Context, err error) {
	switch err {
	case types.ErrInvalidInput, types.ErrDuplicateTitle:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	case types.ErrCreateFailed:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "予期せぬエラーが発生しました。",
		})
	}
}
