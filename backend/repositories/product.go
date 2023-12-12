package repositories

import "e-combomb/models"

var AllProducts []models.Product

func AddProduct(newProduct models.Product) error {
	AllProducts = append(AllProducts, newProduct)

	return nil
}
