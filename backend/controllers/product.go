package controllers

import (
	"e-combomb/models"
	"e-combomb/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type ProductController struct {
	Service *services.ProductService
	store   *sessions.CookieStore
}

func NewProductController(service *services.ProductService, store *sessions.CookieStore) *ProductController {
	return &ProductController{Service: service, store: store}
}

func (pc *ProductController) AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody models.AddProductRequestBody
		session, _ := pc.store.Get(c.Request, "session-name")
		userId, ok := session.Values["user_id"].(uint)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
			return
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request body"})
			return
		}

		if err := pc.Service.AddProduct(requestBody, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "cannot register product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "successfully registered a product"})
	}
}

func (pc *ProductController) GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := pc.Service.GetAllProducts()
		if err != nil {
			fmt.Println("getting products error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": "there's error occurs on the server"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"products": products})
	}
}
