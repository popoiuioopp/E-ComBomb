package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var InMemoryUsers = make(map[string]User)

type User struct {
	Username string
	Password string
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
	user, exists := InMemoryUsers[username]
	fmt.Println(user)
	if !exists {
		return false
	}
	return CheckPasswordHash(username, password)
}
