package routes

import (
	"e-combomb/bootstrap"
	"e-combomb/controllers"
	"e-combomb/middlewares"
	"e-combomb/repositories"
	"e-combomb/services"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func SetupRoutes(router *gin.RouterGroup, store *sessions.CookieStore, app *bootstrap.Application) {

	// Repositories
	productRepository := repositories.NewProductRepository(app.Database)
	authRepository := repositories.NewAuthRepository(app.Database)

	// Services
	productService := services.NewProductService(productRepository)
	authService := services.NewAuthService(authRepository)

	// Middlewares
	authMiddleware := middlewares.NewAuthMiddleWare(store)

	// Controllers
	commonController := controllers.NewCommonController()
	authController := controllers.NewAuthController(authService)
	productController := controllers.NewProductController(productService, store)

	Common(router, commonController)
	AuthRoutes(router, store, authController)
	Product(router, store, productController, authMiddleware)
}

func Common(incomingRoutes *gin.RouterGroup, controller *controllers.CommonController) {
	incomingRoutes.GET("/healthcheck", controller.HealthCheck)
}

func AuthRoutes(router *gin.RouterGroup, store *sessions.CookieStore, controller *controllers.AuthController) {
	router.POST("users/signup", controller.Signup())
	router.POST("users/login", controller.Login(store))
}

func Product(router *gin.RouterGroup, store *sessions.CookieStore, controller *controllers.ProductController, authMiddleware *middlewares.AuthMiddleWare) {
	router.GET("product", controller.GetAllProducts())
	router.POST("product", authMiddleware.SessionAuthMiddleware(store), controller.AddProduct())
}
