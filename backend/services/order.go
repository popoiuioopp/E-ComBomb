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

func (service *OrderService) PlaceOrder(userId uint) error {
	cart, err := service.cartRepo.GetCartByUserId(userId)
	if err != nil {
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
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
		order.Items = append(order.Items, orderItem)
	}

	if err := service.orderRepo.CreateOrder(&order); err != nil {
		return err
	}

	if err := service.cartRepo.ClearCart(cart.Id); err != nil {
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
