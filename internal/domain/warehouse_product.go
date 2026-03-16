package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type WarehouseProduct struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UUID        uuid.UUID `json:"uuid" gorm:"type:uuid;primaryKey"`
	WarehouseID int       `json:"warehouse_id" gorm:"not null"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseID;references:ID"`
	ProductID   int       `json:"product_id" gorm:"not null"`
	Product     Product   `gorm:"foreignKey:ProductID;references:ID"`
	Stock       int       `json:"stock" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	LastSync    time.Time `json:"last_sync"`
}

type WarehouseProductRepository interface {
	Create(ctx context.Context, warehouseProduct *WarehouseProduct) error
	GetByID(ctx context.Context, id uuid.UUID) (*WarehouseProduct, error)
	Fetch(ctx context.Context, limit int, offset int) ([]WarehouseProduct, int64, error)
	Update(ctx context.Context, warehouseProduct *WarehouseProduct) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type WarehouseProductUseCase interface {
	Create(ctx context.Context, warehouseProduct *WarehouseProduct) error
	GetByID(ctx context.Context, id uuid.UUID) (*WarehouseProduct, error)
	Fetch(ctx context.Context, limit int, offset int) ([]WarehouseProduct, int64, error)
	Update(ctx context.Context, warehouseProduct *WarehouseProduct) error
	Delete(ctx context.Context, id uuid.UUID) error
}
