package repositories

import (
	"database/sql"
	"e-combomb/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (repo *OrderRepository) CreateOrder(order *models.Order) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}

	// Insert the order
	orderInsertQuery := "INSERT INTO orders (user_id, created_at, updated_at) VALUES (?, ?, ?)"
	res, err := tx.Exec(orderInsertQuery, order.UserID, order.CreatedAt, order.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert order: %w", err)
	}

	orderID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to retrieve last insert ID for order: %w", err)
	}

	// Insert the order items
	// FIXME: BATCH OPERATION
	for _, item := range order.Items {
		itemInsertQuery := "INSERT INTO order_items (order_id, product_id, quantity, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
		_, err := tx.Exec(itemInsertQuery, orderID, item.ProductID, item.Quantity, item.CreatedAt, item.UpdatedAt)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to insert order item: %w", err)
		}
	}

	return tx.Commit()
}