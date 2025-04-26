package repository

import (
	model "BlogService/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByID(id string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll() ([]model.User, error) {

	var user []model.User

	if err := r.db.Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByID(id string) (*model.User, error) {

	var user model.User

	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
