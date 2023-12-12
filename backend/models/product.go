package models

type Product struct {
	Id          string  `json:"id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	Description string  `json:"description"`
	Owner       string  `json:"owner" binding:"required"`
}

type AddProductRequestBody struct {
	Name        string  `json:"name" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	Description string  `json:"description"`
}
