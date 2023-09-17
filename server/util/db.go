package util

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"job-log/model/entity"
	"log"
)

var DB *gorm.DB

func InitDB() {
	sqlType := viper.GetString("db.sqlType")
	database := viper.GetString("db.database")
	switch sqlType {
	case "mysql":
		username := viper.GetString("db.username")
		password := viper.GetString("db.password")
		host := viper.GetString("db.host")
		port := viper.GetString("db.port")
		dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Panicln("mysql数据库连接失败。", err)
		}
	case "sqlite":
		var err error
		DB, err = gorm.Open(sqlite.Open(database+".db"), &gorm.Config{})
		if err != nil {
			log.Panicln("sqlite数据库连接失败。", err)
		}
	}
}

func InitDBTables() {
	err := DB.AutoMigrate(entity.Todo{})
	if err != nil {
		log.Println("Todo表创建失败")
	}
	err = DB.AutoMigrate(entity.ClipboardLog{})
	if err != nil {
		log.Println("ClipboardLog表创建失败")
	}
}
