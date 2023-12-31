package repositories

import (
	"e-combomb/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) CreateUser(user *models.User) error {
	return ar.db.Create(user).Error
}

func (ar *AuthRepository) GetUser(username string) (models.User, error) {
	var user models.User
	err := ar.db.Where("username = ?", username).Find(&user).Error
	return user, err
}
