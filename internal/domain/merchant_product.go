package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type MerchantProduct struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID        uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique"`
	MerchantID  int       `json:"merchant_id" gorm:"not null"`
	Merchant    Merchant  `gorm:"foreignKey:MerchantID;references:ID"`
	ProductID   int       `json:"product_id" gorm:"not null"`
	Product     Product   `gorm:"foreignKey:ProductID;references:ID"`
	Stock       int       `json:"stock" gorm:"not null"`
	WarehouseID int       `json:"warehouse_id" gorm:"not null"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseID;references:ID"`
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
