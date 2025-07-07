package services

import (
	"airline-system/models"
	"airline-system/repository"
	"airline-system/utils"
	"errors"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) RegisterUser(name, email, password string) (*models.User, error) {
	// Check if email exists
	_, exists := s.UserRepo.FindByEmail(email)
	if exists {
		return nil, errors.New("email already in use")
	}

	// Hash password
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       utils.GenerateID(),
		Name:     name,
		Email:    email,
		Password: hashed,
	}

	err = s.UserRepo.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) LoginUser(email, password string) (string, error) {
	user, exists := s.UserRepo.FindByEmail(email)
	if !exists {
		return "", errors.New("user not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid password")
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
