package main

import (
	"e-combomb/bootstrap"
	"e-combomb/routes"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var (
	store *sessions.CookieStore
)

func init() {
	store = sessions.NewCookieStore([]byte("some-secret-key"))
}

func main() {
	router := gin.New()

	app := bootstrap.App()

	env := app.Env

	gin.SetMode(env.GinMode)

	routes.SetupRoutes(router, store, &app)

	router.Run()
}
