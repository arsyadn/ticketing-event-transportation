package services

import (
	"errors"
	"ticketing-go/models"
	"ticketing-go/repositories"
	"ticketing-go/utils"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user *models.User) (string, uint, error) {
	// hash password
	if err := user.HashPassword(user.Password); err != nil {
		return "", 0, errors.New("failed to hash password")
	}

	// save user
	if err := s.repo.Create(user); err != nil {
		return "", 0, errors.New("failed to create user")
	}

	// generate token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", 0, errors.New("failed to generate token")
	}

	return token, user.ID, nil
}

func (s *UserService) Login(loginRequest *models.LoginRequest) (string, uint, error) {
	user, err := s.repo.FindByEmail(loginRequest.Email)
	if err != nil {
		return "", 0, errors.New("invalid email or password")
	}

	// check password
	if err := user.CheckPassword(loginRequest.Password); err != nil {
		return "", 0, errors.New("invalid email or password")
	}

	// generate token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", 0, errors.New("failed to generate token")
	}

	return token, user.ID, nil
}
