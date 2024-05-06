package models

import (
	"database/sql"
	"time"
)

type Order struct {
	Id        uint        `json:"id"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	UserId    uint        `json:"userId"`
	Items     []OrderItem `json:"items"`
}

type OrderItem struct {
	Id        uint         `json:"id"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"deletedAt"`
	OrderId   uint         `json:"orderId"`
	ProductId uint         `json:"productId"`
	Product   Product      `json:"product"`
	Quantity  int          `json:"quantity"`
}
