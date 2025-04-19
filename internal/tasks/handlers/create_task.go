package handlers

import (
	"go-task-app/internal/tasks/helpers"
	"go-task-app/internal/tasks/repo"
	"go-task-app/internal/tasks/services"
	"go-task-app/internal/tasks/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var taskInput types.TaskHandlerInput

	err := c.ShouldBindJSON(&taskInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "入力値が不正です。",
		})
		return
	}

	task, err := services.CreateTask(
		&types.CreateTaskServiceInput{
			Title:  taskInput.Title,
			UserId: helpers.GetUserIdFromContext(c),
		},
		&services.Repo{
			CreateTask: repo.CreateTask,
		},
	)

	if err != nil {
		helpers.HandleTaskError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"task": task,
	})
}
