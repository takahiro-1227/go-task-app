package handlers

import (
	"go-task-app/internal/tasks/services"
	"go-task-app/internal/tasks/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleTaskError(c *gin.Context, err error) {
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

func CreateTask(c *gin.Context) {
	var newTask types.Task

	err := c.ShouldBindJSON(&newTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "入力値が不正です。",
		})
		return
	}

	task, err := services.CreateTask(newTask)
	if err != nil {
		handleTaskError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}
