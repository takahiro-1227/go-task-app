package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", MysqlUser, MysqlPassword, MysqlHost, MysqlDatabase)), &gorm.Config{})

	if err != nil {
		log.Fatal("データベース接続に失敗しました。")
	}
}
