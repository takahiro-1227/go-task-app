package helpers

import (
	"github.com/gin-gonic/gin"
	"go-task-app/internal/config"
	"go-task-app/internal/tasks/types"
	"log"
	"net/http"
)

func GetUserIdFromContext(c *gin.Context) int {
	userId, exists := c.Get("userId")

	if !exists {
		log.Println("認証が行われていません。")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	return userId.(int)
}

func IsTaskOwnedByUser(taskId uint, userId int) bool {
	var result *types.Task

	config.DB.Select("user_id").Where("id = ?", taskId).Find(&result)

	return userId == result.UserId
}
