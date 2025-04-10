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

var envLoaded bool

func InitIntegrationTest() {
	if !envLoaded {
		config.LoadEnv("../../.env")
		envLoaded = true
	}

	config.ConnectDB()
	initDB()

	if globals.Router == nil {
		globals.Router = routes.SetUpRouter()
	}
}
