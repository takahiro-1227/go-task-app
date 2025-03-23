package handlers

import (
	"go-task-app/internal/config"
	"go-task-app/internal/tasks/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasksHandler(c *gin.Context) {
	var tasks []types.Task

	result := config.DB.Find(&tasks)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "タスクの取得に失敗しました。",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}