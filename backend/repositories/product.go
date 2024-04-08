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
	query := "INSERT INTO products (name, price, description, user_id, product_image) VALUES (?, ?, ?, ?, ?);"
	_, err := pr.db.Exec(query, product.Name, product.Price, product.Description, product.User_id, product.Product_image)
	return err
}

func (pr *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	query := "SELECT id, created_at, updated_at, name, price, description, user_id, product_image FROM products"
	results, err := pr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	for results.Next() {
		var product models.Product
		if err := results.Scan(&product.Id, &product.CreatedAt, &product.UpdatedAt, &product.Name, &product.Price, &product.Description, &product.User_id, &product.Product_image); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
