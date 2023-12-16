package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonController struct {
}

func NewCommonController() *CommonController {
	return &CommonController{}
}

func (cc *CommonController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
