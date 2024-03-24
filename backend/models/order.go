package models

import (
	"database/sql"
	"time"
)

type Order struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	UserID    uint
	Items     []OrderItem
}

type OrderItem struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	OrderID   uint
	ProductID uint
	Product   Product
	Quantity  int
}
