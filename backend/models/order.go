package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint
	Items  []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Product   Product
	Quantity  int
}
