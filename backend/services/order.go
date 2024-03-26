package services

import (
	"e-combomb/models"
	"e-combomb/repositories"
)

type OrderService struct {
	orderRepo *repositories.OrderRepository
	cartRepo  *repositories.CartRepository
}

func NewOrderService(orderRepo *repositories.OrderRepository, cartRepo *repositories.CartRepository) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		cartRepo:  cartRepo,
	}
}

func (service *OrderService) PlaceOrder(userID uint) error {
	cart, err := service.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	order := models.Order{
		UserID: userID,
		Items:  make([]models.OrderItem, 0, len(cart.Items)),
	}

	for _, item := range cart.Items {
		orderItem := models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
		order.Items = append(order.Items, orderItem)
	}

	if err := service.orderRepo.CreateOrder(&order); err != nil {
		return err
	}

	if err := service.cartRepo.ClearCart(cart.ID); err != nil {
		return err
	}

	return nil
}
