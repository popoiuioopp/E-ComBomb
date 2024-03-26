package services

import (
	"e-combomb/models"
	"e-combomb/repositories"
)

type CartService struct {
	repo *repositories.CartRepository
}

func NewCartService(repo *repositories.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (cs *CartService) AddItemToCart(userID uint, item models.CartItem) (models.CartItem, error) {
	cart, err := cs.repo.GetOrCreateCart(userID)
	if err != nil {
		return models.CartItem{}, err
	}

	item.CartID = cart.ID
	return cs.repo.AddItem(item)
}

func (cs *CartService) GetCartByUserID(userID uint) (*models.Cart, error) {
	return cs.repo.GetCartByUserID(userID)
}

func (cs *CartService) RemoveItemFromCart(userID uint, productID uint) error {
	return cs.repo.RemoveItem(userID, productID)
}
