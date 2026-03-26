package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type MerchantResponse struct {
	ID           int       `json:"id"`
	UUID         uuid.UUID `json:"uuid"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Photo        string    `json:"photo"`
	HasWarehouse bool      `json:"has_warehouse"`
	WarehouseId  int       `json:"warehouse_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateMerchantDTO struct {
	Name         string `json:"name" form:"name" validate:"required"`
	Address      string `json:"address" form:"address" validate:"required"`
	Phone        string `json:"phone" form:"phone" validate:"required"`
	Email        string `json:"email" form:"email" validate:"required,email"`
	HasWarehouse bool   `json:"has_warehouse" form:"has_warehouse"`
	WarehouseId  int    `json:"warehouse_id" form:"warehouse_id"`
	KeeperId     int    `json:"keeper_id" form:"keeper_id"`
}

type UpdateMerchantDTO struct {
	UUID         uuid.UUID `json:"uuid" form:"uuid"`
	Name         string    `json:"name" form:"name" validate:"required"`
	Address      string    `json:"address" form:"address" validate:"required"`
	Phone        string    `json:"phone" form:"phone" validate:"required"`
	Email        string    `json:"email" form:"email" validate:"required,email"`
	HasWarehouse bool      `json:"has_warehouse" form:"has_warehouse"`
	WarehouseId  int       `json:"warehouse_id" form:"warehouse_id"`
	KeeperId     int       `json:"keeper_id" form:"keeper_id"`
}

func FromMerchant(merchant domain.Merchant) MerchantResponse {
	return MerchantResponse{
		ID:           merchant.ID,
		UUID:         merchant.UUID,
		Name:         merchant.Name,
		Address:      merchant.Address,
		Phone:        merchant.Phone,
		Email:        merchant.Email,
		Photo:        merchant.Photo,
		HasWarehouse: merchant.HasWarehouse,
		WarehouseId:  merchant.WarehouseId,
		CreatedAt:    merchant.CreatedAt,
	}
}

func FromMerchants(merchants []domain.Merchant) []MerchantResponse {
	merchantResponses := make([]MerchantResponse, 0)
	for _, merchant := range merchants {
		merchantResponses = append(merchantResponses, FromMerchant(merchant))
	}
	return merchantResponses
}
