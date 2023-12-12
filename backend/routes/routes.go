package routes

import (
	"e-combomb/controllers"
	"e-combomb/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func SetupRoutes(router *gin.Engine, store *sessions.CookieStore) {
	Common(router)
	AuthRoutes(router, store)
	Product(router, store)
}

func Common(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/healthcheck", controllers.HealthCheck)
}

func AuthRoutes(router *gin.Engine, store *sessions.CookieStore) {
	router.POST("users/signup", controllers.Signup())
	router.POST("users/login", controllers.Login(store))
}

func Product(router *gin.Engine, store *sessions.CookieStore) {
	router.GET("product", controllers.GetAllProducts())
	router.POST("product", middlewares.SessionAuthMiddleware(store), controllers.AddProduct())
}
