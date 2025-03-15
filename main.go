package main

import (
	"go-task-app/internal/config"
	"go-task-app/internal/tasks"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	router := gin.Default()

	router.GET("/tasks", tasks.GetTasks)
	router.POST("/task", tasks.CreateTask)
	err := router.Run("localhost:4000")

	if err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
