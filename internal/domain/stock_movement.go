package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type StockMovement struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID         uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique"`
	WarehouseID  int       `json:"warehouse_id" gorm:"not null"`
	Warehouse    Warehouse `gorm:"foreignKey:WarehouseID;references:ID"`
	MerchantID   int       `json:"merchant_id" gorm:"not null"`
	Merchant     Merchant  `gorm:"foreignKey:MerchantID;references:ID"`
	ProductID    int       `json:"product_id" gorm:"not null"`
	Product      Product   `gorm:"foreignKey:ProductID;references:ID"`
	Type         string    `json:"type" gorm:"not null;type:varchar(30)"`
	ReferenceId  int       `json:"reference_id" gorm:"not null;"`
	Quantity     int       `json:"quantity" gorm:"not null"`
	StrockBefore int       `json:"strock_before" gorm:"not null"`
	StrockAfter  int       `json:"strock_after" gorm:"not null"`
	Reason       string    `json:"reason" gorm:"not null;type:varchar(255)"`
	TenantID     int       `json:"tenant_id" gorm:"not null"`
	Tenant       Tenant    `gorm:"foreignKey:TenantID;references:ID"`
	CreatedBy    string    `json:"created_by" gorm:"not null;type:varchar(255)"`
	UpdatedBy    string    `json:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy    string    `json:"deleted_by" gorm:"type:varchar(255)"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	LastSync     time.Time `json:"last_sync"`
}

type StockMovementRepository interface {
	Create(ctx context.Context, stockMovement *StockMovement) error
	GetByID(ctx context.Context, id uuid.UUID) (*StockMovement, error)
	Fetch(ctx context.Context, limit int, offset int) ([]StockMovement, int64, error)
	Update(ctx context.Context, stockMovement *StockMovement) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type StockMovementUseCase interface {
	Create(ctx context.Context, stockMovement *StockMovement) error
	GetByID(ctx context.Context, id uuid.UUID) (*StockMovement, error)
	Fetch(ctx context.Context, limit int, offset int) ([]StockMovement, int64, error)
	Update(ctx context.Context, stockMovement *StockMovement) error
	Delete(ctx context.Context, id uuid.UUID) error
}
