package helpers

import (
	"go-task-app/internal/config"
	"go-task-app/internal/routes"
	"go-task-app/tests/globals"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func initDB() {
	migration, err := migrate.New(
		"file://../../migrations",
		"mysql://"+config.MysqlUser+":"+config.MysqlPassword+"@tcp("+config.MysqlHost+")/"+config.MysqlDatabase,
	)

	if err != nil {
		panic(err)
	}

	err = migration.Down()

	if err != nil {
		panic(err)
	}

	err = migration.Up()

	if err != nil && err.Error() != "no change" {
		panic(err)
	}
}

func InitTest() {
	config.LoadEnv("../../.env")
	initDB()
	config.ConnectDB()

	globals.Router = routes.SetUpRouter()
}
