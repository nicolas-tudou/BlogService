package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId   string `gorm:"unique;not null"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gotm:"not null"`
}
