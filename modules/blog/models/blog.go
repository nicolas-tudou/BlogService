package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	Auth    User   `gorm:"foreignKey:UserId"`
	UserId  string
}
