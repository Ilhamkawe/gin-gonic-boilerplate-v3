package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type WarehouseProduct struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID        uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique"`
	WarehouseID int       `json:"warehouse_id" gorm:"not null"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseID;references:ID"`
	ProductID   int       `json:"product_id" gorm:"not null"`
	Product     Product   `gorm:"foreignKey:ProductID;references:ID"`
	Stock       int       `json:"stock" gorm:"not null"`
	TenantID    int       `json:"tenant_id" gorm:"not null"`
	Tenant      Tenant    `gorm:"foreignKey:TenantID;references:ID"`
	CreatedBy   string    `json:"created_by" gorm:"not null"`
	UpdatedBy   string    `json:"updated_by" gorm:""`
	DeletedBy   string    `json:"deleted_by" gorm:""`
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
