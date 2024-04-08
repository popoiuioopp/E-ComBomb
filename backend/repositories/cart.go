package repositories

import (
	"database/sql"
	"e-combomb/models"

	_ "github.com/go-sql-driver/mysql"
)

type CartRepository struct {
	DB *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{DB: db}
}

func (cr *CartRepository) GetOrCreateCart(userId uint) (*models.Cart, error) {
	var cart models.Cart
	query := "SELECT id, created_at, updated_at, deleted_at, user_id FROM carts WHERE user_id = ?"
	err := cr.DB.QueryRow(query, userId).Scan(&cart.Id, &cart.CreatedAt, &cart.UpdatedAt, &cart.DeletedAt, &cart.UserId)

	switch {
	case err == sql.ErrNoRows:
		// Cart does not exist, create a new one
		insertQuery := "INSERT INTO carts (user_id) VALUES (?)"
		result, err := cr.DB.Exec(insertQuery, userId)
		if err != nil {
			return nil, err
		}

		cartId, err := result.LastInsertId()
		if err != nil {
			return nil, err
		}

		cart.Id = uint(cartId)
		cart.UserId = userId

	case err != nil:
		return nil, err
	}

	return &cart, nil
}

func (cr *CartRepository) AddItem(cartItem models.CartItem) (models.CartItem, error) {
	var existingItem models.CartItem
	query := "SELECT id, cart_id, product_id, quantity FROM cart_items WHERE cart_id = ? AND product_id = ?"
	err := cr.DB.QueryRow(query, cartItem.CartId, cartItem.ProductId).Scan(&existingItem.Id, &existingItem.CartId, &existingItem.ProductId, &existingItem.Quantity)

	if err == sql.ErrNoRows {
		// Create a new cart item
		insertQuery := "INSERT INTO cart_items (cart_id, product_id, quantity) VALUES (?, ?, ?)"
		_, err := cr.DB.Exec(insertQuery, cartItem.CartId, cartItem.ProductId, cartItem.Quantity)
		if err != nil {
			return models.CartItem{}, err
		}
		return cartItem, nil
	} else if err != nil {
		return models.CartItem{}, err
	}

	// Update existing item quantity
	existingItem.Quantity += cartItem.Quantity
	updateQuery := "UPDATE cart_items SET quantity = ? WHERE id = ?"
	_, err = cr.DB.Exec(updateQuery, existingItem.Quantity, existingItem.Id)
	if err != nil {
		return models.CartItem{}, err
	}

	return existingItem, nil
}

func (cr *CartRepository) GetCartByUserId(userId uint) (*models.Cart, error) {
	var cart models.Cart
	cartQuery := "SELECT id, created_at, updated_at, deleted_at, user_id FROM carts WHERE user_id = ?"
	err := cr.DB.QueryRow(cartQuery, userId).Scan(&cart.Id, &cart.CreatedAt, &cart.UpdatedAt, &cart.DeletedAt, &cart.UserId)
	if err != nil {
		return nil, err
	}

	itemsQuery := `
	SELECT ci.id, ci.cart_id, ci.product_id, ci.quantity, ci.created_at, ci.updated_at, ci.deleted_at, 
	       p.name, p.price, p.description 
	FROM cart_items ci
	JOIN products p ON ci.product_id = p.id
	WHERE ci.cart_id = ?`

	rows, err := cr.DB.Query(itemsQuery, cart.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cart.Items = []models.CartItem{}
	for rows.Next() {
		var item models.CartItem
		var product models.Product
		if err := rows.Scan(&item.Id, &item.CartId, &item.ProductId, &item.Quantity, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt, &product.Name, &product.Price, &product.Description); err != nil {
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

func (cr *CartRepository) ClearCart(cartId uint) error {
	query := "DELETE FROM cart_items WHERE cart_id = ?"
	_, err := cr.DB.Exec(query, cartId)
	return err
}

func (cr *CartRepository) RemoveItem(userId uint, productId uint) error {
	_, err := cr.DB.Exec("DELETE FROM cart_items WHERE cart_id = (SELECT id FROM carts WHERE user_id = ?) AND product_id = ?", userId, productId)
	return err
}

func (cr *CartRepository) UpdateItemQuantity(cartId uint, productId uint, quantity int) error {
	query := "UPDATE cart_items SET quantity = ? WHERE cart_id = ? AND product_id = ?"
	_, err := cr.DB.Exec(query, quantity, cartId, productId)
	return err
}
