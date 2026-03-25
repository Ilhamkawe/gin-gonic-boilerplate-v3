package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type ProductResponse struct {
	ID            int        `json:"id"`
	UUID          uuid.UUID  `json:"uuid"`
	Name          string     `json:"name"`
	Thumbnail     string     `json:"thumbnail"`
	Description   string     `json:"description"`
	Price         int        `json:"price"`
	CategoryID    int        `json:"category_id"`
	IsPopular     bool       `json:"is_popular"`
	AttributeJson string     `json:"attribute_json"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

type CreateProductDTO struct {
	Name          string `json:"name" form:"name" validate:"required"`
	Description   string `json:"description" form:"description" validate:"required"`
	Price         int    `json:"price" form:"price" validate:"required"`
	CategoryID    int    `json:"category_id" form:"category_id" validate:"required"`
	IsPopular     bool   `json:"is_popular" form:"is_popular"`
	AttributeJson string `json:"attribute_json" form:"attribute_json"`
}

type UpdateProductDTO struct {
	UUID          uuid.UUID `json:"uuid" form:"uuid"`
	Name          string    `json:"name" form:"name" validate:"required"`
	Description   string    `json:"description" form:"description" validate:"required"`
	Price         int       `json:"price" form:"price" validate:"required"`
	CategoryID    int       `json:"category_id" form:"category_id" validate:"required"`
	IsPopular     bool      `json:"is_popular" form:"is_popular"`
	AttributeJson string    `json:"attribute_json" form:"attribute_json"`
}

func FromProduct(product domain.Product) ProductResponse {
	return ProductResponse{
		ID:            product.ID,
		UUID:          product.UUID,
		Name:          product.Name,
		Thumbnail:     product.Thumbnail,
		Description:   product.Description,
		Price:         product.Price,
		CategoryID:    product.CategoryID,
		IsPopular:     product.IsPopular,
		AttributeJson: product.AttributeJson,
		CreatedAt:     product.CreatedAt,
		UpdatedAt:     product.UpdatedAt,
	}
}

func FromProducts(products []domain.Product) []ProductResponse {
	productResponses := make([]ProductResponse, 0)
	for _, product := range products {
		productResponses = append(productResponses, FromProduct(product))
	}
	return productResponses
}
