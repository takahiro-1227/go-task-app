package helpers

import (
	"go-task-app/internal/config"
	"go-task-app/internal/routes"
	"go-task-app/tests/globals"
)

func initDB() {
	config.DB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	config.DB.Exec("TRUNCATE TABLE users")
	config.DB.Exec("TRUNCATE TABLE tasks")
	config.DB.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func InitTest() {
	config.LoadEnv("../../.env")
	config.ConnectDB()
	initDB()

	globals.Router = routes.SetUpRouter()
}
