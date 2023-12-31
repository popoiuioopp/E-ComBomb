package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	Description string  `json:"description"`
	User_id     uint    `json:"user_id" binding:"required"`
}

type AddProductRequestBody struct {
	Name        string  `json:"name" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	Description string  `json:"description"`
}

type ProductInterface struct {
	ID          uint    `json:"id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	Description string  `json:"description"`
	User_id     uint    `json:"user_id" binding:"required"`
}

func MapProductsToProductInterfaces(products []Product) []ProductInterface {
	var productInterfaces []ProductInterface
	for _, product := range products {
		productInterface := ProductInterface{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			User_id:     product.User_id,
		}
		productInterfaces = append(productInterfaces, productInterface)
	}
	return productInterfaces
}
