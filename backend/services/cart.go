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

func (cs *CartService) AddItemToCart(userId uint, item models.CartItem) (models.CartItem, error) {
	cart, err := cs.repo.GetOrCreateCart(userId)
	if err != nil {
		return models.CartItem{}, err
	}

	item.CartId = cart.Id
	return cs.repo.AddItem(item)
}

func (cs *CartService) GetCartByUserId(userId uint) (*models.Cart, error) {
	return cs.repo.GetCartByUserId(userId)
}

func (cs *CartService) RemoveItemFromCart(userId uint, productId uint) error {
	return cs.repo.RemoveItem(userId, productId)
}

func (cs *CartService) UpdateItemQuantity(userId uint, productId uint, quantity int) error {
	// Get the user's cart
	cart, err := cs.repo.GetCartByUserId(userId)
	if err != nil {
		return err
	}

	// Update the quantity of the specified item
	return cs.repo.UpdateItemQuantity(cart.Id, productId, quantity)
}
