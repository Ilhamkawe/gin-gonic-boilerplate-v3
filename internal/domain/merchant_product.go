package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type MerchantProduct struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UUID        uuid.UUID `json:"uuid" gorm:"type:uuid;primaryKey"`
	MerchantID  int       `json:"merchant_id" gorm:"not null"`
	Merchant    Merchant  `gorm:"foreignKey:MerchantID;references:ID"`
	ProductID   int       `json:"product_id" gorm:"not null"`
	Product     Product   `gorm:"foreignKey:ProductID;references:ID"`
	Stock       int       `json:"stock" gorm:"not null"`
	WarehouseID int       `json:"warehouse_id" gorm:"not null"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseID;references:ID"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	LastSync    time.Time `json:"last_sync"`
}

type MerchantProductRepository interface {
	Create(ctx context.Context, merchantProduct *MerchantProduct) error
	GetByID(ctx context.Context, id uuid.UUID) (*MerchantProduct, error)
	Fetch(ctx context.Context, limit int, offset int) ([]MerchantProduct, int64, error)
	Update(ctx context.Context, merchantProduct *MerchantProduct) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type MerchantProductUseCase interface {
	Create(ctx context.Context, merchantProduct *MerchantProduct) error
	GetByID(ctx context.Context, id uuid.UUID) (*MerchantProduct, error)
	Fetch(ctx context.Context, limit int, offset int) ([]MerchantProduct, int64, error)
	Update(ctx context.Context, merchantProduct *MerchantProduct) error
	Delete(ctx context.Context, id uuid.UUID) error
}
