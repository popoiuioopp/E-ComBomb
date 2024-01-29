package main

import (
	"e-combomb/bootstrap"
	"e-combomb/routes"
	"time"

	"github.com/gin-contrib/cors"
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

	// Custom CORS configuration
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))

	apiRoutes := router.Group("/api")

	app := bootstrap.App()
	env := app.Env
	gin.SetMode(env.GinMode)

	routes.SetupRoutes(apiRoutes, store, &app)

	router.Run()
}
