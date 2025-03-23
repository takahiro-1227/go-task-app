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

	router.GET("/tasks", tasks_handlers.GetTasksHandler)
	router.POST("/task", tasks_handlers.CreateTaskHandler)

	err := router.Run(":3000")

	if err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
