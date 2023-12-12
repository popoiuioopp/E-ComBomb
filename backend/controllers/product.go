package controllers

import (
	"e-combomb/models"
	"e-combomb/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody models.AddProductRequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request body"})
			return
		}

		if err := services.AddProduct(requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "cannot register product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "successfully registered a product"})
	}
}

func GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"products": services.GetAllProducts()})
	}
}
