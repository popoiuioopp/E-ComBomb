package main

import (
	"e-combomb/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/healthcheck", controller.HealthCheck)

	r.Run()
}
