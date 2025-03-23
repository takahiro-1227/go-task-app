package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlUser string
var MysqlPassword string
var MysqlHost string
var MysqlDatabase string
var Host string

func init() {
	env := os.Getenv("GO_ENV")
	if env == "development" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal(".env を読み込めませんでした。")
		}
	}

	MysqlUser = os.Getenv("MYSQL_USER")
	MysqlPassword = os.Getenv("MYSQL_PASSWORD")
	MysqlHost = os.Getenv("MYSQL_HOST")
	MysqlDatabase = os.Getenv("MYSQL_DATABASE")
	Host = os.Getenv("HOST")
}

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", MysqlUser, MysqlPassword, MysqlHost, MysqlDatabase)), &gorm.Config{})

	if err != nil {
		log.Fatal("データベース接続に失敗しました。")
	}
}
