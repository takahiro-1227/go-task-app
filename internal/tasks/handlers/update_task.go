package handlers

import (
	"go-task-app/internal/tasks/helpers"
	"go-task-app/internal/tasks/services"
	"go-task-app/internal/tasks/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "IDは数値である必要があります。",
		})
		return
	}

	var newTask types.Task

	err = c.ShouldBindJSON(&newTask)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "入力値が不正です。",
		})
		return
	}

	newTask.ID = uint(id)

	result, err := services.UpdateTask(&newTask)

	if err != nil {
		helpers.HandleTaskError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": result,
	})
}
