package bootstrap

import (
	"database/sql"
	"e-combomb/models"
	"fmt"
)

type Database interface {
	Find(out interface{}, where ...interface{}) error
	CreateUser(user *models.User) error
	ValidateUserCredentials(username, password string) (bool, *models.User, error)
}

func NewMySQLDatabase(env *Env) *sql.DB {
	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)

	if err != nil {
		panic(err.Error())
	}

	return db
}
