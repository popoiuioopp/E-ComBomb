package models

type Cart struct {
	Id     uint
	UserId uint
	Items  []CartItem
}

type CartItem struct {
	Id        uint
	CartId    uint
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
