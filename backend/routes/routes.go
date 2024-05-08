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
	cartRepository := repositories.NewCartRepository(app.Database)
	orderRepository := repositories.NewOrderRepository(app.Database)

	// Services
	productService := services.NewProductService(productRepository)
	authService := services.NewAuthService(authRepository)
	cartService := services.NewCartService(cartRepository)
	orderService := services.NewOrderService(orderRepository, cartRepository)
	paymentService := services.NewPaymentService(orderRepository)

	// Middleware
	authMiddleware := middlewares.NewAuthMiddleWare(store)

	// Controllers
	commonController := controllers.NewCommonController()
	authController := controllers.NewAuthController(authService)
	productController := controllers.NewProductController(productService, store)
	cartController := controllers.NewCartController(cartService, store)
	orderController := controllers.NewOrderController(orderService, store)
	paymentController := controllers.NewPaymentController(paymentService, store)

	Common(router, commonController)
	AuthRoutes(router, store, authController)
	Product(router, store, productController, authMiddleware)
	Cart(router, store, cartController, authMiddleware)
	Order(router, store, orderController, authMiddleware)
	Payment(router, store, paymentController, authMiddleware)
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

func Cart(router *gin.RouterGroup, store *sessions.CookieStore, controller *controllers.CartController, authMiddleware *middlewares.AuthMiddleWare) {
	router.GET("cart", authMiddleware.SessionAuthMiddleware(store), controller.GetCart())
	router.POST("cart", authMiddleware.SessionAuthMiddleware(store), controller.AddItemToCart())
	router.DELETE("/cart/:productId", authMiddleware.SessionAuthMiddleware(store), controller.RemoveItemFromCart())
	router.PUT("/cart/:productId", authMiddleware.SessionAuthMiddleware(store), controller.UpdateCartItem())
}

func Order(router *gin.RouterGroup, store *sessions.CookieStore, controller *controllers.OrderController, authMiddleware *middlewares.AuthMiddleWare) {
	router.POST("order", authMiddleware.SessionAuthMiddleware(store), controller.PlaceOrder())
	router.GET("order", authMiddleware.SessionAuthMiddleware(store), controller.GetOrders())
	router.GET("order/:orderId", authMiddleware.SessionAuthMiddleware(store), controller.GetOrderById())
}

func Payment(router *gin.RouterGroup, store *sessions.CookieStore, controller *controllers.PaymentController, authMiddleware *middlewares.AuthMiddleWare) {
	router.POST("payment/checkout/:orderId", authMiddleware.SessionAuthMiddleware(store), controller.CheckoutOrder())
}
