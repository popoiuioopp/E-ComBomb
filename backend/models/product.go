package models

type Product struct {
	Id           uint    `json:"id"`
	Name         string  `json:"name" binding:"required"`
	Price        float32 `json:"price" binding:"required"`
	Description  string  `json:"description"`
	UserId       uint    `json:"user_id" binding:"required"`
	ProductImage string  `json:"product_image"`
}

type AddProductRequestBody struct {
	Name         string  `json:"name" binding:"required"`
	Price        float32 `json:"price" binding:"required"`
	Description  string  `json:"description"`
	ProductImage string  `json:"product_image"`
}

type ProductInterface struct {
	Id           uint    `json:"id" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	Price        float32 `json:"price" binding:"required"`
	Description  string  `json:"description"`
	UserId       uint    `json:"user_id" binding:"required"`
	ProductImage string  `json:"product_image"`
}

func MapProductsToProductInterfaces(products []Product) []ProductInterface {
	var productInterfaces []ProductInterface
	for _, product := range products {
		productInterface := ProductInterface{
			Id:           product.Id,
			Name:         product.Name,
			Price:        product.Price,
			Description:  product.Description,
			UserId:       product.UserId,
			ProductImage: product.ProductImage,
		}
		productInterfaces = append(productInterfaces, productInterface)
	}
	return productInterfaces
}
