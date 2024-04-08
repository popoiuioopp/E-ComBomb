package controllers

import (
	"e-combomb/models"
	"e-combomb/services"
	"net/http"
	"strconv"

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
		var requestCartItem models.AddToCartRequest
		if err := c.ShouldBindJSON(&requestCartItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cartItem := models.CartItem{ProductID: requestCartItem.ProductId, Quantity: requestCartItem.Quantity}

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

func (cc *CartController) RemoveItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := cc.store.Get(c.Request, "session-name")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get session"})
			return
		}

		userID, ok := session.Values["user_id"].(uint)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
			return
		}

		productID, err := strconv.ParseUint(c.Param("productId"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
			return
		}

		err = cc.service.RemoveItemFromCart(userID, uint(productID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove item from cart"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "item removed from cart"})
	}
}

func (cc *CartController) UpdateCartItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := cc.store.Get(c.Request, "session-name")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get session"})
			return
		}

		userID, ok := session.Values["user_id"].(uint)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
			return
		}

		// Extracting product ID and new quantity from the request
		productID, _ := strconv.ParseUint(c.Param("productId"), 10, 32)
		var json models.UpdateCartRequest
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Updating the item in the cart
		err = cc.service.UpdateItemQuantity(userID, uint(productID), json.Quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item quantity"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Item quantity updated successfully"})
	}
}
