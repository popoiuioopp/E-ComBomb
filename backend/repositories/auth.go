package repositories

import (
	"database/sql"
	"e-combomb/models"

	_ "github.com/go-sql-driver/mysql"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?);"
	_, err := ar.db.Exec(query, user.Username, user.Password)
	return err
}

func (ar *AuthRepository) GetUser(username string) (models.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = ?;"
	var user models.User
	err := ar.db.QueryRow(query, username).Scan(&user.Id, &user.Username, &user.Password)
	return user, err
}
