package handlers

import (
	"go-task-app/internal/tasks/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "IDは数値である必要があります。",
		})
		return
	}

	serviceErr := services.DeleteTask(id)
	if serviceErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "タスクを削除できませんでした。",
		})
		return
	}

	c.String(200, "Success")
}
