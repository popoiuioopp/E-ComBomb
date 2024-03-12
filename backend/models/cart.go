package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId uint
	Items  []CartItem
}

type CartItem struct {
	gorm.Model
	CartID    uint
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}
