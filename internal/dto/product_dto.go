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
	UpdatedAt     *time.Time               `json:"updated_at"`
	Variants      []ProductVariantResponse `json:"variants,omitempty"`
}

type CreateProductDTO struct {
	Name          string                    `json:"name" validate:"required"`
	Description   string                    `json:"description" validate:"required"`
	Price         int                       `json:"price" validate:"required"`
	CategoryID    int                       `json:"category_id" validate:"required"`
	IsPopular     bool                      `json:"is_popular"`
	Thumbnail     string                    `json:"thumbnail"`
	AttributeJson string                    `json:"attribute_json"`
	Variants      []CreateProductVariantDTO `json:"variants"`
}

type UpdateProductDTO struct {
	UUID          uuid.UUID `json:"uuid"`
	Name          string    `json:"name" validate:"required"`
	Description   string    `json:"description" validate:"required"`
	Price         int       `json:"price" validate:"required"`
	CategoryID    int       `json:"category_id" validate:"required"`
	IsPopular     bool      `json:"is_popular"`
	Thumbnail     string    `json:"thumbnail"`
	AttributeJson string    `json:"attribute_json"`
}

func FromProduct(product domain.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		UUID:        product.UUID,
		Name:        product.Name,
		Thumbnail:   product.Thumbnail,
		Description: product.Description,
		CategoryID:  product.CategoryID,
		IsPopular:   product.IsPopular,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		Variants:    FromProductVariants(product.Variants),
	}
}

func FromProducts(products []domain.Product) []ProductResponse {
	productResponses := make([]ProductResponse, 0)
	for _, product := range products {
		productResponses = append(productResponses, FromProduct(product))
	}
	return productResponses
}
