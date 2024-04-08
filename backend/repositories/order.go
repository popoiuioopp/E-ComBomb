package repositories

import (
	"database/sql"
	"e-combomb/models"
	"fmt"
	"strings"

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
	orderInsertQuery := "INSERT INTO orders (user_id) VALUES (?)"
	res, err := tx.Exec(orderInsertQuery, order.UserId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert order: %w", err)
	}

	orderId, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to retrieve last insert ID for order: %w", err)
	}

	// Prepare batch insert for order items
	valueStrings := []string{}
	valueArgs := []interface{}{}
	for _, item := range order.Items {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, orderId, item.ProductId, item.Quantity, item.CreatedAt, item.UpdatedAt)
	}

	itemInsertQuery := fmt.Sprintf("INSERT INTO order_items (order_id, product_id, quantity, created_at, updated_at) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err = tx.Exec(itemInsertQuery, valueArgs...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert order items: %w", err)
	}

	return tx.Commit()
}
