package controllers

import (
	"e-combomb/services"
	"fmt"
	"net/http"
	"strconv"

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
		userId, ok := session.Values["user_id"].(uint)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
			return
		}

		if err := oc.orderService.PlaceOrder(userId); err != nil {
			fmt.Println("Error placing order: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not place order"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "order placed successfully"})
	}
}

func (oc *OrderController) GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := oc.store.Get(c.Request, "session-name")
		userId, ok := session.Values["user_id"].(uint)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
			return
		}

		Orders, err := oc.orderService.GetAllOrders(int(userId))
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "there's an error occurs"})
		}

		c.JSON(http.StatusOK, gin.H{"orders": Orders})
	}
}

func (oc *OrderController) GetOrderById() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := oc.store.Get(c.Request, "session-name")
		userId, ok := session.Values["user_id"].(uint)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
			return
		}

		orderId, err := strconv.Atoi(c.Param("orderId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
			return
		}

		order, err := oc.orderService.GetOrderById(int(userId), orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting order"})
			return
		}

		if order == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
			return
		}

		c.JSON(http.StatusOK, order)
	}
}
