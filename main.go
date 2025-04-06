package main

import (
	"go-task-app/internal/config"
	"go-task-app/internal/routes"
	"log"
)

func main() {
	config.LoadEnv(".env")
	config.ConnectDB()

	router := routes.SetUpRouter()

	err := router.Run(":4000")

	if err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
