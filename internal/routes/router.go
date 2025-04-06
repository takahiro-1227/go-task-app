package routes

import (
	"go-task-app/internal/middlewares"
	tasks_handlers "go-task-app/internal/tasks/handlers"
	users_handlers "go-task-app/internal/users/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
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

	return router
}
