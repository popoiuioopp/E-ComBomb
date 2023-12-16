package main

import (
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

	routes.SetupRoutes(router, store)

	router.Run()
}
