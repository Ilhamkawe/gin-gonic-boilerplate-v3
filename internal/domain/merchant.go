package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Merchant struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID         uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique"`
	Name         string    `json:"name" gorm:"not null"`
	Address      string    `json:"address" gorm:"not null"`
	Phone        string    `json:"phone" gorm:"not null"`
	Email        string    `json:"email" gorm:"not null"`
	Photo        string    `json:"photo" gorm:"not null"`
	KeeperId     int       `json:"keeper_id" gorm:"not null"`
	Keeper       User      `gorm:"foreignKey:KeeperId;references:ID"`
	HasWarehouse bool      `json:"has_warehouse" gorm:"not null"`
	WarehouseId  int       `json:"warehouse_id" gorm:"not null"`
	Warehouse    Warehouse `gorm:"foreignKey:WarehouseId;references:ID"`
	TenantID     int       `json:"tenant_id" gorm:"not null"`
	Tenant       Tenant    `gorm:"foreignKey:TenantID;references:ID"`
	CreatedBy    string    `json:"created_by" gorm:"not null"`
	UpdatedBy    string    `json:"updated_by" gorm:""`
	DeletedBy    string    `json:"deleted_by" gorm:""`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	LastSync     time.Time `json:"last_sync"`
}

type MerchantRepository interface {
	Create(ctx context.Context, merchant *Merchant) error
	GetByID(ctx context.Context, id uuid.UUID) (*Merchant, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Merchant, int64, error)
	Update(ctx context.Context, merchant *Merchant) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type MerchantUseCase interface {
	Create(ctx context.Context, merchant *Merchant) error
	GetByID(ctx context.Context, id uuid.UUID) (*Merchant, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Merchant, int64, error)
	Update(ctx context.Context, merchant *Merchant) error
	Delete(ctx context.Context, id uuid.UUID) error
}
