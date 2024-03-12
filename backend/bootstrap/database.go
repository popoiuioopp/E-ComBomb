package bootstrap

import (
	"e-combomb/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
	Find(out interface{}, where ...interface{}) error
	CreateUser(user *models.User) error
	ValidateUserCredentials(username, password string) (bool, *models.User, error)
}

func NewMySQLDatabase(env *Env) *gorm.DB {

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println(mysqlURI)
	db, err := gorm.Open(mysql.Open(mysqlURI), &gorm.Config{})

	if err != nil {
		panic("Cannot instantiate database's connection")
	}

	// AutoMigrate with slice of models
	models := []interface{}{
		&models.Product{},
		&models.User{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	}

	err = db.AutoMigrate(models...)
	if err != nil {
		log.Fatal("failed to auto-migrate:", err)
	}

	return db
}
