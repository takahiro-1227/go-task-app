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
		authGroup.GET("/", func(c *gin.Context) {
			userId, exists := c.Get("userId")
			if !exists {
				c.Status(401)
				return
			}
			c.JSON(200, gin.H{
				"status": "ok",
				"userId": userId,
			})
		})
	}

	router.GET("/tasks", tasks_handlers.GetTasks)
	router.POST("/task", tasks_handlers.CreateTask)
	router.PUT("/task/:id", tasks_handlers.UpdateTask)
	router.DELETE("/task/:id", tasks_handlers.DeleteTask)

	router.POST("/sign-up", users_handlers.SignUp)
	router.POST("/sign-in", users_handlers.SignIn)

	err := router.Run(":4000")

	if err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
