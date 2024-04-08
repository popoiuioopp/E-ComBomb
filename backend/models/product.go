package models

import (
	"database/sql"
	"time"
)

type Product struct {
	Id            uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime
	Name          string  `json:"name" binding:"required"`
	Price         float32 `json:"price" binding:"required"`
	Description   string  `json:"description"`
	User_id       uint    `json:"user_id" binding:"required"`
	Product_image string  `json:"product_image"`
}

type AddProductRequestBody struct {
	Name          string  `json:"name" binding:"required"`
	Price         float32 `json:"price" binding:"required"`
	Description   string  `json:"description"`
	Product_image string  `json:"product_image"`
}

type ProductInterface struct {
	Id            uint    `json:"id" binding:"required"`
	Name          string  `json:"name" binding:"required"`
	Price         float32 `json:"price" binding:"required"`
	Description   string  `json:"description"`
	User_id       uint    `json:"user_id" binding:"required"`
	Product_image string  `json:"product_image"`
}

func MapProductsToProductInterfaces(products []Product) []ProductInterface {
	var productInterfaces []ProductInterface
	for _, product := range products {
		productInterface := ProductInterface{
			Id:            product.Id,
			Name:          product.Name,
			Price:         product.Price,
			Description:   product.Description,
			User_id:       product.User_id,
			Product_image: product.Product_image,
		}
		productInterfaces = append(productInterfaces, productInterface)
	}
	return productInterfaces
}
