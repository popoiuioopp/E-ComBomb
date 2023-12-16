package services

import (
	"e-combomb/models"
	"e-combomb/repositories"

	"github.com/google/uuid"
)

type ProductService struct {
	productRepository *repositories.ProductRepositry
}

func NewProductService(productRepository *repositories.ProductRepositry) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (s *ProductService) AddProduct(productRequestBody models.AddProductRequestBody) error {
	newProduct := models.Product{
		Id:          uuid.NewString(),
		Name:        productRequestBody.Name,
		Description: productRequestBody.Description,
		Price:       productRequestBody.Price,
	}

	return s.productRepository.AddProduct(newProduct)
}

func (s *ProductService) GetAllProducts() []models.Product {
	return s.productRepository.GetAllProducts()
}
