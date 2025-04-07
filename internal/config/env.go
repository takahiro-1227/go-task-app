package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var MysqlUser string
var MysqlPassword string
var MysqlHost string
var MysqlDatabase string
var Host string
var AuthSecret string

func LoadEnv(envPath string) {
	env := os.Getenv("GO_ENV")
	if env == "development" || env == "testing" {
		err := godotenv.Load(envPath)

		if err != nil {
			log.Println(err)
			log.Fatal(".env を読み込めませんでした。")
		}
	}

	MysqlUser = os.Getenv("MYSQL_USER")
	MysqlPassword = os.Getenv("MYSQL_PASSWORD")
	MysqlHost = os.Getenv("MYSQL_HOST")
	MysqlDatabase = os.Getenv("MYSQL_DATABASE")
	Host = os.Getenv("HOST")
	AuthSecret = os.Getenv("AUTH_SECRET")

	if env == "testing" {
		MysqlDatabase = MysqlDatabase + "_testing"
	}
}
