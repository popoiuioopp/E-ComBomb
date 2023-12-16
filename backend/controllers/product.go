package controllers

import (
	"e-combomb/models"
	"e-combomb/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	Service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{Service: service}
}

func (pc *ProductController) AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody models.AddProductRequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request body"})
			return
		}

		if err := pc.Service.AddProduct(requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "cannot register product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "successfully registered a product"})
	}
}

func (pc *ProductController) GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"products": pc.Service.GetAllProducts()})
	}
}
