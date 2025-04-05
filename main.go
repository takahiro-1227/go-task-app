package main

import (
	"go-task-app/internal/config"
	"go-task-app/internal/middlewares"
	tasks_handlers "go-task-app/internal/tasks/handlers"
	users_handlers "go-task-app/internal/users/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.InitDB()

	router := gin.Default()

	authGroup := router.Group("/")

	authGroup.Use(middlewares.Auth())
	{
		authGroup.GET("/tasks", tasks_handlers.GetTasks)
		authGroup.POST("/task", tasks_handlers.CreateTask)
		authGroup.PUT("/task/:id", tasks_handlers.UpdateTask)
		authGroup.DELETE("/task/:id", tasks_handlers.DeleteTask)
	}
	router.POST("/sign-up", users_handlers.SignUp)
	router.POST("/sign-in", users_handlers.SignIn)

	err := router.Run(":4000")

	if err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
