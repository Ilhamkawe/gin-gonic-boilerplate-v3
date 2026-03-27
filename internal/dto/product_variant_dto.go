package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"gorm.io/datatypes"
)

type ProductVariantResponse struct {
	ID            int            `json:"id"`
	UUID          uuid.UUID      `json:"uuid"`
	ProductID     int            `json:"product_id"`
	ProductName   string         `json:"product_name,omitempty"`
	SKU           string         `json:"sku"`
	AttributeJson datatypes.JSON `json:"attribute_json"`
	Price         float64        `json:"price"`
	Barcode       string         `json:"barcode"`
	ImageUrl      string         `json:"image"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
}

type CreateProductVariantDTO struct {
	ProductID     int            `json:"product_id" form:"product_id" validate:"required"`
	SKU           string         `json:"sku" form:"sku" validate:"required"`
	AttributeJson datatypes.JSON `json:"attribute_json" form:"attribute_json"`
	Price         float64        `json:"price" form:"price" validate:"required"`
	Barcode       string         `json:"barcode" form:"barcode"`
}

type UpdateProductVariantDTO struct {
	SKU           string         `json:"sku" form:"sku" validate:"required"`
	AttributeJson datatypes.JSON `json:"attribute_json" form:"attribute_json"`
	Price         float64        `json:"price" form:"price" validate:"required"`
	Barcode       string         `json:"barcode" form:"barcode"`
}

func FromProductVariant(variant domain.ProductVariant) ProductVariantResponse {
	return ProductVariantResponse{
		ID:            variant.ID,
		UUID:          variant.UUID,
		ProductID:     variant.ProductID,
		ProductName:   variant.Product.Name,
		SKU:           variant.SKU,
		AttributeJson: variant.AttributeJson,
		Price:         variant.Price,
		Barcode:       variant.Barcode,
		CreatedAt:     variant.CreatedAt,
		UpdatedAt:     variant.UpdatedAt,
	}
}

func FromProductVariants(variants []domain.ProductVariant) []ProductVariantResponse {
	variantResponses := make([]ProductVariantResponse, 0)
	for _, variant := range variants {
		variantResponses = append(variantResponses, FromProductVariant(variant))
	}
	return variantResponses
}
