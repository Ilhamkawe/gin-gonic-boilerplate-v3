package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type WarehouseResponse struct {
	ID        int       `json:"id"`
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateWarehouseDTO struct {
	Name    string `json:"name" form:"name" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
	Phone   string `json:"phone" form:"phone" validate:"required"`
	Email   string `json:"email" form:"email" validate:"required,email"`
}

type UpdateWarehouseDTO struct {
	UUID    uuid.UUID `json:"uuid" form:"uuid"`
	Name    string    `json:"name" form:"name" validate:"required"`
	Address string    `json:"address" form:"address" validate:"required"`
	Phone   string    `json:"phone" form:"phone" validate:"required"`
	Email   string    `json:"email" form:"email" validate:"required,email"`
}

func FromWarehouse(warehouse domain.Warehouse) WarehouseResponse {
	return WarehouseResponse{
		ID:        warehouse.ID,
		UUID:      warehouse.UUID,
		Name:      warehouse.Name,
		Address:   warehouse.Address,
		Phone:     warehouse.Phone,
		Email:     warehouse.Email,
		Photo:     warehouse.Photo,
		CreatedAt: warehouse.CreatedAt,
	}
}

func FromWarehouses(warehouses []domain.Warehouse) []WarehouseResponse {
	warehouseResponses := make([]WarehouseResponse, 0)
	for _, warehouse := range warehouses {
		warehouseResponses = append(warehouseResponses, FromWarehouse(warehouse))
	}
	return warehouseResponses
}
