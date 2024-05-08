package controllers

import (
	"e-combomb/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type PaymentController struct {
	store   *sessions.CookieStore
	service *services.PaymentService
}

func NewPaymentController(service *services.PaymentService, store *sessions.CookieStore) *PaymentController {
	return &PaymentController{
		store:   store,
		service: service,
	}
}

func (pc *PaymentController) CheckoutOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := pc.store.Get(c.Request, "session-name")
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

		err = pc.service.CheckoutOrder(userId, orderId, "completed")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update order status"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "order completed successfully"})

	}
}
