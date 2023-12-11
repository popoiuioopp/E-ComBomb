package routes

import (
	"e-combomb/controllers"
	controller "e-combomb/controllers"
	middleware "e-combomb/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func SetupRoutes(router *gin.Engine, store *sessions.CookieStore) {
	AuthRoutes(router, store)
	Common(router)
}

func AuthRoutes(router *gin.Engine, store *sessions.CookieStore) {
	router.Use(middleware.SessionAuth(store))
	router.POST("users/signup", controllers.Signup())
	router.POST("users/login", controllers.Login(store))
}

func Common(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/healthcheck", controller.HealthCheck)
}
