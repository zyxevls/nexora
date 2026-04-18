package service

import (
	"errors"
	"nexora/internal/model"
	"nexora/internal/repository"
	"nexora/internal/utils"
)

func Register(name, email, password string) (*model.User, error) {
	// Check if user already exists
	existingUser, err := repository.FindUserByEmail(email)

	if err == nil && existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create the user
	user := &model.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	err = repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func Login(email, password string) (*model.User, error) {
	// Find the user by email
	user, err := repository.FindUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Verify the password
	if !utils.CheckPassword(password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

func ValidateToken(token string) (*model.User, error) {
	userID, err := utils.ParseToken(token)
	if err != nil {
		return nil, err
	}
	user, err := repository.FindUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
