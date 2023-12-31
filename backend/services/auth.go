package services

import (
	"e-combomb/models"
	"e-combomb/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepository *repositories.AuthRepository
}

func NewAuthService(authRepository *repositories.AuthRepository) *AuthService {
	return &AuthService{authRepository: authRepository}
}

func (as *AuthService) CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return as.authRepository.CreateUser(user)
}

func (as *AuthService) ValidateUserCredentials(username, password string) (bool, models.User, error) {
	user, err := as.authRepository.GetUser(username)
	if err != nil {
		return false, models.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false, models.User{}, err
	}

	return true, user, nil
}
