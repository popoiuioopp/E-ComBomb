package models

import (
	"database/sql"
	"time"
)

type Order struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	UserId    uint
	Items     []OrderItem
}

type OrderItem struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	OrderId   uint
	ProductId uint
	Product   Product
	Quantity  int
}
