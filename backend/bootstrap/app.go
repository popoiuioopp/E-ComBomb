package bootstrap

import (
	"database/sql"
)

type Application struct {
	Env      *Env
	Database *sql.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Database = NewMySQLDatabase(app.Env)
	return *app
}
