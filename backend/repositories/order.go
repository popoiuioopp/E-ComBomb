package repositories

import (
	"database/sql"
	"e-combomb/models"
	"errors"
	"fmt"
	"strings"
	"time"

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
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, orderId, item.ProductId, item.Quantity)
	}

	itemInsertQuery := fmt.Sprintf("INSERT INTO order_items (order_id, product_id, quantity) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err = tx.Exec(itemInsertQuery, valueArgs...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert order items: %w", err)
	}

	return tx.Commit()
}

func (repo *OrderRepository) GetAllOrders(userId int) ([]models.Order, error) {
	var orders []models.Order
	var orderMap = make(map[uint]*models.Order)

	query := `SELECT 
                o.id AS order_id,
                o.created_at AS order_created_at,
                o.updated_at AS order_updated_at,
                o.deleted_at AS order_deleted_at,
                o.user_id AS order_user_id,
				o.status AS order_status,
                oi.id AS item_id,
                oi.order_id AS item_order_id,
                oi.product_id AS item_product_id,
                oi.quantity AS item_quantity,
                p.id AS product_id,
                p.name AS product_name,
                p.price AS product_price,
                p.description AS product_description,
                p.user_id AS product_user_id,
                p.product_image AS product_image
            FROM
                orders o
            INNER JOIN
                order_items oi ON o.id = oi.order_id
            INNER JOIN
                products p ON oi.product_id = p.id
            WHERE
                o.user_id = ?;`

	results, err := repo.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	for results.Next() {
		var orderItem models.OrderItem
		var product models.Product
		var orderId, orderUserId uint
		var orderCreatedAt, orderUpdatedAt time.Time
		var orderDeletedAt sql.NullTime
		var orderStatus string

		err := results.Scan(
			&orderId,
			&orderCreatedAt,
			&orderUpdatedAt,
			&orderDeletedAt,
			&orderUserId,
			&orderStatus,
			&orderItem.Id,
			&orderItem.OrderId,
			&orderItem.ProductId,
			&orderItem.Quantity,
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Description,
			&product.UserId,
			&product.ProductImage,
		)
		if err != nil {
			return nil, err
		}

		orderItem.Product = product

		if existingOrder, exists := orderMap[orderId]; exists {
			existingOrder.Items = append(existingOrder.Items, orderItem)
		} else {
			orderMap[orderId] = &models.Order{
				Id:        orderId,
				CreatedAt: orderCreatedAt,
				UpdatedAt: orderUpdatedAt,
				UserId:    orderUserId,
				Items:     []models.OrderItem{orderItem},
			}
		}
	}

	for _, order := range orderMap {
		orders = append(orders, *order)
	}

	return orders, nil
}

func (repo *OrderRepository) GetOrderById(userId, orderId int) (*models.Order, error) {
	orderMap := make(map[uint]*models.Order)

	query := `SELECT 
                o.id AS order_id,
                o.created_at AS order_created_at,
                o.updated_at AS order_updated_at,
                o.deleted_at AS order_deleted_at,
                o.user_id AS order_user_id,
				o.status AS order_status,
                oi.id AS item_id,
                oi.order_id AS item_order_id,
                oi.product_id AS item_product_id,
                oi.quantity AS item_quantity,
                p.id AS product_id,
                p.name AS product_name,
                p.price AS product_price,
                p.description AS product_description,
                p.user_id AS product_user_id,
                p.product_image AS product_image
            FROM
                orders o
            INNER JOIN
                order_items oi ON o.id = oi.order_id
            INNER JOIN
                products p ON oi.product_id = p.id
            WHERE
                o.user_id = ? AND o.id = ?;`

	rows, err := repo.DB.Query(query, userId, orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var orderItem models.OrderItem
		var product models.Product
		var orderId, orderUserId uint
		var orderCreatedAt, orderUpdatedAt time.Time
		var orderDeletedAt sql.NullTime
		var orderStatus string

		err = rows.Scan(
			&orderId,
			&orderCreatedAt,
			&orderUpdatedAt,
			&orderDeletedAt,
			&orderUserId,
			&orderStatus,
			&orderItem.Id,
			&orderItem.OrderId,
			&orderItem.ProductId,
			&orderItem.Quantity,
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Description,
			&product.UserId,
			&product.ProductImage,
		)
		if err != nil {
			return nil, err
		}

		orderItem.Product = product

		if order, exists := orderMap[orderId]; exists {
			order.Items = append(order.Items, orderItem)
		} else {
			orderMap[orderId] = &models.Order{
				Id:        orderId,
				CreatedAt: orderCreatedAt,
				UpdatedAt: orderUpdatedAt,
				UserId:    orderUserId,
				Items:     []models.OrderItem{orderItem},
				Status:    orderStatus,
			}
		}
	}

	// If the order is found, it will be the only key in the map
	for _, order := range orderMap {
		return order, nil
	}

	// If no order was found, return nil
	return nil, errors.New("order not found")
}
