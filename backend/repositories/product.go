package repositories

import (
	"e-combomb/models"

	"gorm.io/gorm"
)

type ProductRepositry struct {
	DB *gorm.DB
}

func NewProductRepository(database *gorm.DB) *ProductRepositry {
	return &ProductRepositry{DB: database}
}

func (pr *ProductRepositry) NewProductRepository() *ProductRepositry {
	return &ProductRepositry{}
}

func (pr *ProductRepositry) AddProduct(product *models.Product) error {
	return pr.DB.Create(product).Error
}

func (pr *ProductRepositry) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := pr.DB.Find(&products)
	return products, result.Error
}
