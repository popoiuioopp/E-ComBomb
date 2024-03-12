package controllers

import (
	"e-combomb/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type OrderController struct {
	orderService *services.OrderService
	store        *sessions.CookieStore
}

func NewOrderController(orderService *services.OrderService, store *sessions.CookieStore) *OrderController {
	return &OrderController{
		orderService: orderService,
		store:        store,
	}
}

func (oc *OrderController) PlaceOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := oc.store.Get(c.Request, "session-name")
		userID, ok := session.Values["user_id"].(uint)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
			return
		}

		if err := oc.orderService.PlaceOrder(userID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not place order"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "order placed successfully"})
	}
}
