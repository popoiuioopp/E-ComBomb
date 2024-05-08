package services

import (
	"e-combomb/repositories"
)

type PaymentService struct {
	orderRepo *repositories.OrderRepository
}

func NewPaymentService(orderRepo *repositories.OrderRepository) *PaymentService {
	return &PaymentService{
		orderRepo: orderRepo,
	}
}

func (service *PaymentService) CheckoutOrder(userId uint, orderId int, status string) error {
	return service.orderRepo.UpdateOrderStatus(userId, orderId, status)
}
