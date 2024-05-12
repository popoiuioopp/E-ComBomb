package models

type Cart struct {
	Id     uint       `json:"id"`
	UserId uint       `json:"userId"`
	Items  []CartItem `json:"items"`
}

type CartItem struct {
	Id        uint    `json:"id"`
	CartId    uint    `json:"cartId"`
	ProductId uint    `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required"`
	Product   Product `json:"product"`
}

type AddToCartRequest struct {
	ProductId uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type UpdateCartRequest struct {
	Quantity int `json:"quantity"`
}
