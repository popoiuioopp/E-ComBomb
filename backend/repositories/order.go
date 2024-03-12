package repositories

import (
	"e-combomb/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (repo *OrderRepository) CreateOrder(order *models.Order) error {
	return repo.DB.Create(order).Error
}
