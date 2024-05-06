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
	err := cr.DB.QueryRow("CALL GetOrCreateCart_v1(?)", userId).Scan(&cart.Id, &cart.UserId)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (cr *CartRepository) AddItem(cartItem models.CartItem) (models.CartItem, error) {
	_, err := cr.DB.Exec("CALL AddItem_v1(?, ?, ?)", cartItem.CartId, cartItem.ProductId, cartItem.Quantity)
	if err != nil {
		return models.CartItem{}, err
	}
	return cartItem, nil
}

func (cr *CartRepository) GetCartByUserId(userId uint) (models.Cart, error) {
	var cart models.Cart

	// Call the stored procedure
	rows, err := cr.DB.Query("CALL GetCartByUserId_v1(?)", userId)
	if err != nil {
		return models.Cart{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&cart.Id, &cart.UserId)
		if err != nil {
			return models.Cart{}, err
		}
	}

	// Go to next result set
	rows.NextResultSet()

	for rows.Next() {
		var cartItem models.CartItem

		err := rows.Scan(&cartItem.Id, &cartItem.ProductId, &cartItem.Quantity, &cartItem.Product.Name,
			&cartItem.Product.Price, &cartItem.Product.Description)
		if err != nil {
			return models.Cart{}, err
		}
		cart.Items = append(cart.Items, cartItem)
	}

	return cart, nil
}

func (cr *CartRepository) ClearCart(cartId uint) error {
	_, err := cr.DB.Exec("CALL ClearCart_v1(?)", cartId)
	return err
}

func (cr *CartRepository) RemoveItem(userId uint, productId uint) error {
	_, err := cr.DB.Exec("CALL RemoveItem_v1(?, ?)", userId, productId)
	return err
}

func (cr *CartRepository) UpdateItemQuantity(cartId uint, productId uint, quantity int) error {
	_, err := cr.DB.Exec("CALL UpdateItemQuantity_v1(?, ?, ?)", cartId, productId, quantity)
	return err
}
