package db

import (
	config "BlogService/config"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	once.Do(func() {
		appConfig := config.AppConfig
		var err error
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", appConfig.DBUser, appConfig.DBPassword, appConfig.DBHost, appConfig.DBPort, appConfig.DBName)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("数据库连接失败：%v, dsn: %s", err, dsn)
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("数据库连接失败")
		}

		if err = sqlDB.Ping(); err != nil {
			log.Fatal("数据库连接失败，无法ping")
		}

		fmt.Printf("数据库连接成功")
	})
	return db
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	sqlDB.Close()
}
