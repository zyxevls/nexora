package repository

import (
	"nexora/config"
	"nexora/internal/model"
)

func CreateUser(user *model.User) error {
	return config.DB.Create(user).Error
}

func FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func FindUserByID(id uint) (*model.User, error) {
	var user model.User
	err := config.DB.First(&user, id).Error
	return &user, err
}
