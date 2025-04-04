package handlers

import (
	"go-task-app/internal/tasks/helpers"
	"go-task-app/internal/tasks/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks, err := services.GetTasks()

	if err != nil {
		helpers.HandleTaskError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}
