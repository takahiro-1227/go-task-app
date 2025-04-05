package handlers

import (
	"go-task-app/internal/tasks/helpers"
	"go-task-app/internal/tasks/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	userId := helpers.GetUserIdFromContext(c)

	tasks, err := services.GetTasks(userId)

	if err != nil {
		helpers.HandleTaskError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}
