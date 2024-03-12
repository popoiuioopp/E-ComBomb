package services

import (
	"e-combomb/models"
	"e-combomb/repositories"
)

type ProductService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (s *ProductService) AddProduct(productRequestBody models.AddProductRequestBody, userId uint) error {
	newProduct := models.Product{
		Name:        productRequestBody.Name,
		Description: productRequestBody.Description,
		Price:       productRequestBody.Price,
		User_id:     userId,
	}

	return s.productRepository.AddProduct(&newProduct)
}

func (s *ProductService) GetAllProducts() ([]models.ProductInterface, error) {
	products, err := s.productRepository.GetAllProducts()
	if err != nil {
		return []models.ProductInterface{}, err
	}

	productInterfaces := models.MapProductsToProductInterfaces(products)

	return productInterfaces, nil
}
