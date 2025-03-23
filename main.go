package main

import (
	"go-task-app/internal/config"
	tasks_handlers "go-task-app/internal/tasks/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	router := gin.Default()

	router.GET("/tasks", tasks_handlers.GetTasks)
	router.POST("/task", tasks_handlers.CreateTask)
	router.DELETE("/task/:id", tasks_handlers.DeleteTask)

	err := router.Run(":4000")

	if err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
