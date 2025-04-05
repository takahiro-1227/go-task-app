package handlers

import (
	"go-task-app/internal/tasks/helpers"
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

	userId := helpers.GetUserIdFromContext(c)

	err = services.DeleteTask(uint(id), userId)
	if err != nil {
		helpers.HandleTaskError(c, err)
		return
	}

	c.String(200, "タスクを削除しました。")
}
