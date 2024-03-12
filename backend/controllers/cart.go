package controllers

import (
	"e-combomb/models"
	"e-combomb/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type CartController struct {
	service *services.CartService
	store   *sessions.CookieStore
}

func NewCartController(service *services.CartService, store *sessions.CookieStore) *CartController {
	return &CartController{service: service, store: store}
}

func (cc *CartController) AddItemToCart() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Retrieve the session
		session, err := cc.store.Get(c.Request, "session-name")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get session"})
			return
		}

		// Extract the userID from the session
		userID, ok := session.Values["user_id"].(uint)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
			return
		}

		// Bind the incoming JSON payload to the cartItem structure
		var cartItem models.CartItem
		if err := c.ShouldBindJSON(&cartItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Add item to cart using the service
		result, err := cc.service.AddItemToCart(userID, cartItem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func (cc *CartController) GetCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := cc.store.Get(c.Request, "session-name")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve session"})
			return
		}

		userID, ok := session.Values["user_id"].(uint)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		cart, err := cc.service.GetCartByUserID(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart"})
			return
		}

		c.JSON(http.StatusOK, cart)
	}
}
