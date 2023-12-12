package services

import (
	"e-combomb/models"
	"e-combomb/repositories"

	"github.com/google/uuid"
)

func AddProduct(productRequestBody models.AddProductRequestBody) error {
	newproduct := models.Product{
		Id:          uuid.New().String(),
		Name:        productRequestBody.Name,
		Description: productRequestBody.Description,
		Price:       productRequestBody.Price,
	}

	return repositories.AddProduct(newproduct)
}

func GetAllProducts() []models.Product {
	return repositories.AllProducts
}
