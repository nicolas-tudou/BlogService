package db

import (
	"sync"

	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetBlogDB() *gorm.DB {
	once.Do(func() {
		db = getDB("blog")
	})
	return db
}

func CloseBlogDB() {
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	sqlDB.Close()
}
