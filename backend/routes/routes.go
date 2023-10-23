package routes

import (
	controller "craft-cart/controllers"
	middleware "craft-cart/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.UserAuthenticate())
	incomingRoutes.GET("/usersdata", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}

func Common(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/healthcheck", controller.HealthCheck)
}