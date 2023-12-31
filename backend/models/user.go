package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var InMemoryUsers = make(map[string]User)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func CheckPasswordHash(username, password string) bool {
	hashedPassword := InMemoryUsers[username].Password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

func CheckPassword(username, password string) bool {
	_, exists := InMemoryUsers[username]
	if !exists {
		return false
	}
	return CheckPasswordHash(username, password)
}
