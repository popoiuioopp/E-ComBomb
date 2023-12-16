package repositories

import "e-combomb/models"

var AllProducts []models.Product

type ProductRepositry struct {
}

func NewProductRepository() *ProductRepositry {
	return &ProductRepositry{}
}

func (pr *ProductRepositry) NewProductRepository() *ProductRepositry {
	return &ProductRepositry{}
}

func (pr *ProductRepositry) AddProduct(newProduct models.Product) error {
	AllProducts = append(AllProducts, newProduct)

	return nil
}

func (pr *ProductRepositry) GetAllProducts() []models.Product {
	return AllProducts
}
