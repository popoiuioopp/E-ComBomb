package models

import (
	"database/sql"
	"time"
)

type Cart struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	UserId    uint
	Items     []CartItem
}

type CartItem struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	CartID    uint
	ProductID uint    `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required"`
	Product   Product `json:"product"`
}

type AddToCartRequest struct {
	ProductId uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
