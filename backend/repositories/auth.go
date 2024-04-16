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
	_, err := ar.db.Exec("CALL CreateUser_v1(?, ?)", user.Username, user.Password)
	return err
}

func (ar *AuthRepository) GetUser(username string) (models.User, error) {
	var user models.User
	err := ar.db.QueryRow("CALL GetUser_v1(?)", username).Scan(&user.Id, &user.Username, &user.Password)
	return user, err
}
