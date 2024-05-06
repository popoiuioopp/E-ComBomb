package models

import (
	"time"
)

type Order struct {
	Id        uint        `json:"id"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	UserId    uint        `json:"userId"`
	Items     []OrderItem `json:"items"`
}

type OrderItem struct {
	Id        uint    `json:"id"`
	OrderId   uint    `json:"orderId"`
	ProductId uint    `json:"productId"`
	Product   Product `json:"product"`
	Quantity  int     `json:"quantity"`
}
