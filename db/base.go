package db

import (
	config "BlogService/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDB(name string) *gorm.DB {
	appConfig := config.AppConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", appConfig.DBUser, appConfig.DBPassword, appConfig.DBHost, appConfig.DBPort, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败：%v, dsn: %s", err, dsn)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("数据库连接失败")
	}

	if err = sqlDB.Ping(); err != nil {
		log.Fatal("数据库%s连接失败，无法ping", name)
	}

	fmt.Printf("数据库%s连接成功", name)

	return db
}
