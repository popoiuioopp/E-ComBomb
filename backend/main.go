package main

import (
	"craft-cart/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	routes.Common(router)
	routes.UserRoutes(router)
    routes.AuthRoutes(router)

	router.Run()
}
