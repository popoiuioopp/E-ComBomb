package repositories

import (
	"e-combomb/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(database *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: database}
}

func (pr *ProductRepository) NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (pr *ProductRepository) AddProduct(product *models.Product) error {
	return pr.DB.Create(product).Error
}

func (pr *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := pr.DB.Find(&products)
	return products, result.Error
}
