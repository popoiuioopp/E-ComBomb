package controllers

import (
	"e-combomb/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Products []models.Product

func AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newProduct models.Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "bad request body"})
		}
		Products = append(Products, newProduct)

		c.JSON(http.StatusOK, gin.H{"message": "successfully registered a product"})
	}
}

func GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"products": Products})
	}
}
