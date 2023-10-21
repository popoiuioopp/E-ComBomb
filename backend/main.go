package main

import (
	"craft-cart/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/healthcheck", controller.HealthCheck)

	r.Run()
}
