package repositories

import (
	"e-combomb/models"
	"errors"

	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{DB: db}
}

func (cr *CartRepository) GetOrCreateCart(userID uint) (*models.Cart, error) {
	var cart models.Cart
	if err := cr.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Cart does not exist, create a new one
			cart = models.Cart{UserId: userID}
			if err := cr.DB.Create(&cart).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return &cart, nil
}

func (cr *CartRepository) AddItem(cartItem models.CartItem) (models.CartItem, error) {
	var existingItem models.CartItem
	if err := cr.DB.Where("cart_id = ? AND product_id = ?", cartItem.CartID, cartItem.ProductID).First(&existingItem).Error; err == nil {
		// Update existing item quantity
		existingItem.Quantity += cartItem.Quantity
		cr.DB.Save(&existingItem)
		return existingItem, nil
	}

	// Create a new cart item
	if err := cr.DB.Create(&cartItem).Error; err != nil {
		return models.CartItem{}, err
	}

	return cartItem, nil
}

func (cr *CartRepository) GetCartByUserID(userID uint) (*models.Cart, error) {
	var cart models.Cart
	if err := cr.DB.Where("user_id = ?", userID).Preload("Items").First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

func (cr *CartRepository) ClearCart(cartID uint) error {
	return cr.DB.Where("cart_id = ?", cartID).Delete(&models.CartItem{}).Error
}
