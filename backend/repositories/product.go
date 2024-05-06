package repositories

import (
	"database/sql"
	"e-combomb/models"

	_ "github.com/go-sql-driver/mysql"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(database *sql.DB) *ProductRepository {
	return &ProductRepository{db: database}
}

func (pr *ProductRepository) AddProduct(product *models.Product) error {
	_, err := pr.db.Exec("CALL AddProduct_v1(?, ?, ?, ?, ?)", product.Name, product.Price, product.Description, product.UserId, product.ProductImage)
	return err
}

func (pr *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	results, err := pr.db.Query("CALL GetAllProducts_v1()")
	if err != nil {
		return nil, err
	}
	defer results.Close()

	for results.Next() {
		var product models.Product
		if err := results.Scan(&product.Id, &product.Name, &product.Price, &product.Description, &product.UserId, &product.ProductImage); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
