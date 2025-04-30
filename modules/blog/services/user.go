package services

import (
	model "BlogService/modules/blog/models"
	repository "BlogService/modules/blog/repositories"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id string) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) GetAllUsers() ([]model.User, error) {

	return u.userRepo.FindAll()
}

func (u *userService) GetUserByID(id string) (*model.User, error) {

	return u.userRepo.FindByID(id)
}
