package routes

import (
	"e-combomb/controllers"
	"e-combomb/middlewares"
	"e-combomb/repositories"
	"e-combomb/services"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func SetupRoutes(router *gin.Engine, store *sessions.CookieStore) {

	// Repositories
	productRepository := repositories.NewProductRepository()

	// Services
	productService := services.NewProductService(productRepository)

	// Middlewares
	authMiddleware := middlewares.NewAuthMiddleWare(store)

	// Controllers
	commonController := controllers.NewCommonController()
	authController := controllers.NewAuthController()
	productController := controllers.NewProductController(productService)

	Common(router, commonController)
	AuthRoutes(router, store, authController)
	Product(router, store, productController, authMiddleware)
}

func Common(incomingRoutes *gin.Engine, controller *controllers.CommonController) {
	incomingRoutes.GET("/healthcheck", controller.HealthCheck)
}

func AuthRoutes(router *gin.Engine, store *sessions.CookieStore, controller *controllers.AuthController) {
	router.POST("users/signup", controller.Signup())
	router.POST("users/login", controller.Login(store))
}

func Product(router *gin.Engine, store *sessions.CookieStore, controller *controllers.ProductController, authMiddleware *middlewares.AuthMiddleWare) {
	router.GET("product", controller.GetAllProducts())
	router.POST("product", authMiddleware.SessionAuthMiddleware(store), controller.AddProduct())
}
