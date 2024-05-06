package services

import (
	"e-combomb/models"
	"e-combomb/repositories"
	"log"
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

func (service *OrderService) PlaceOrder(userId uint) error {
	cart, err := service.cartRepo.GetCartByUserId(userId)
	if err != nil {
		log.Printf("Error getting cart by user id: %v\n", err)
		return err
	}

	order := models.Order{
		UserId: userId,
		Items:  make([]models.OrderItem, 0, len(cart.Items)),
	}

	for _, item := range cart.Items {
		orderItem := models.OrderItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		}
		order.Items = append(order.Items, orderItem)
	}

	if err := service.orderRepo.CreateOrder(&order); err != nil {
		log.Printf("Error creating Order: %v\n", err)
		return err
	}

	if err := service.cartRepo.ClearCart(cart.Id); err != nil {
		log.Printf("Error clearing cart: %v\n", err)
		return err
	}

	return nil
}

func (service *OrderService) GetAllOrders(userId int) ([]models.Order, error) {
	return service.orderRepo.GetAllOrders(userId)
}

func (service *OrderService) GetOrderById(userId, orderId int) (*models.Order, error) {
	return service.orderRepo.GetOrderById(userId, orderId)
}
