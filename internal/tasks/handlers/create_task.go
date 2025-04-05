package handlers

import (
	"go-task-app/internal/tasks/helpers"
	"go-task-app/internal/tasks/services"
	"go-task-app/internal/tasks/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var newTask types.Task

	err := c.ShouldBindJSON(&newTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "入力値が不正です。",
		})
		return
	}

	newTask.UserId = helpers.GetUserIdFromContext(c)

	task, err := services.CreateTask(&newTask)
	if err != nil {
		helpers.HandleTaskError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}
