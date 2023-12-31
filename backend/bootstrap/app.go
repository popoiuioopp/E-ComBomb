package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env      *Env
	Database *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Database = NewMySQLDatabase(app.Env)
	return *app
}
