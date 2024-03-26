package repositories

import (
	"database/sql"
	"e-combomb/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type CartRepository struct {
	DB *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{DB: db}
}

func (cr *CartRepository) GetOrCreateCart(userID uint) (*models.Cart, error) {
	var cart models.Cart
	query := "SELECT id, created_at, updated_at, deleted_at, user_id FROM carts WHERE user_id = ?"
	err := cr.DB.QueryRow(query, userID).Scan(&cart.ID, &cart.CreatedAt, &cart.UpdatedAt, &cart.DeletedAt, &cart.UserId)

	switch {
	case err == sql.ErrNoRows:
		// Cart does not exist, create a new one
		insertQuery := "INSERT INTO carts (user_id) VALUES (?)"
		result, err := cr.DB.Exec(insertQuery, userID)
		if err != nil {
			return nil, err
		}

		cartID, err := result.LastInsertId()
		if err != nil {
			return nil, err
		}

		cart.ID = uint(cartID)
		cart.UserId = userID

	case err != nil:
		return nil, err
	}

	return &cart, nil
}

func (cr *CartRepository) AddItem(cartItem models.CartItem) (models.CartItem, error) {
	var existingItem models.CartItem
	query := "SELECT id, cart_id, product_id, quantity FROM cart_items WHERE cart_id = ? AND product_id = ?"
	err := cr.DB.QueryRow(query, cartItem.CartID, cartItem.ProductID).Scan(&existingItem.ID, &existingItem.CartID, &existingItem.ProductID, &existingItem.Quantity)

	if err == sql.ErrNoRows {
		// Create a new cart item
		insertQuery := "INSERT INTO cart_items (cart_id, product_id, quantity) VALUES (?, ?, ?)"
		_, err := cr.DB.Exec(insertQuery, cartItem.CartID, cartItem.ProductID, cartItem.Quantity)
		if err != nil {
			fmt.Println("HelloWorld1")
			return models.CartItem{}, err
		}
		return cartItem, nil
	} else if err != nil {
		fmt.Println("HelloWorld2")
		return models.CartItem{}, err
	}

	// Update existing item quantity
	existingItem.Quantity += cartItem.Quantity
	updateQuery := "UPDATE cart_items SET quantity = ? WHERE id = ?"
	_, err = cr.DB.Exec(updateQuery, existingItem.Quantity, existingItem.ID)
	if err != nil {
		fmt.Println("HelloWorld3")
		return models.CartItem{}, err
	}

	return existingItem, nil
}

func (cr *CartRepository) GetCartByUserID(userID uint) (*models.Cart, error) {
	var cart models.Cart
	cartQuery := "SELECT id, created_at, updated_at, deleted_at, user_id FROM carts WHERE user_id = ?"
	err := cr.DB.QueryRow(cartQuery, userID).Scan(&cart.ID, &cart.CreatedAt, &cart.UpdatedAt, &cart.DeletedAt, &cart.UserId)
	if err != nil {
		return nil, err
	}

	itemsQuery := `
	SELECT ci.id, ci.cart_id, ci.product_id, ci.quantity, ci.created_at, ci.updated_at, ci.deleted_at, 
	       p.name, p.price, p.description 
	FROM cart_items ci
	JOIN products p ON ci.product_id = p.id
	WHERE ci.cart_id = ?`

	rows, err := cr.DB.Query(itemsQuery, cart.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cart.Items = []models.CartItem{}
	for rows.Next() {
		var item models.CartItem
		var product models.Product
		if err := rows.Scan(&item.ID, &item.CartID, &item.ProductID, &item.Quantity, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt, &product.Name, &product.Price, &product.Description); err != nil {
			return nil, err
		}
		item.Product = product
		cart.Items = append(cart.Items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &cart, nil
}

func (cr *CartRepository) ClearCart(cartID uint) error {
	query := "DELETE FROM cart_items WHERE cart_id = ?"
	_, err := cr.DB.Exec(query, cartID)
	return err
}

func (cr *CartRepository) RemoveItem(userID uint, productID uint) error {
	_, err := cr.DB.Exec("DELETE FROM cart_items WHERE cart_id = (SELECT id FROM carts WHERE user_id = ?) AND product_id = ?", userID, productID)
	return err
}
