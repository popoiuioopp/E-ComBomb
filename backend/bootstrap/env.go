package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv  string `mapstructure:"APP_ENV"`
	DBHost  string `mapstructure:"DB_HOST"`
	DBPort  string `mapstructure:"DB_PORT"`
	DBUser  string `mapstructure:"DB_USER"`
	DBPass  string `mapstructure:"DB_PASS"`
	DBName  string `mapstructure:"DB_NAME"`
	GinMode string `mapstructure:"GIN_MODE"`
}

func NewEnv() *Env {
	viper.AutomaticEnv() // Read from environment variables first

	viper.SetConfigFile(".env")
	err := viper.MergeInConfig() // Merge in the values from .env file, if present
	if err != nil {
		log.Printf("Warning: No .env file found or error in reading it: %v\n", err)
	}

	env := Env{}
	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
