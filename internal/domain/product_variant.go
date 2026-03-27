package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ProductVariant struct {
	ID            int            `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID          uuid.UUID      `json:"uuid" gorm:"type:uuid;not null;unique;default:gen_random_uuid()"`
	ProductID     int            `json:"product_id" gorm:"not null"`
	Product       Product        `gorm:"foreignKey:ProductID;references:ID"`
	SKU           string         `json:"sku" gorm:"not null;unique"`
	AttributeJson datatypes.JSON `json:"attribute_json" gorm:"type:jsonb"`
	TenantID      int            `json:"tenant_id" gorm:"not null;index"`
	Tenant        Tenant         `gorm:"foreignKey:TenantID;references:ID"`
	Price         float64        `json:"price" gorm:"not null"`
	Barcode       string         `json:"barcode" gorm:"unique"`
	CreatedBy     string         `json:"created_by" gorm:"type:varchar(255);not null"`
	UpdatedBy     string         `json:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy     string         `json:"deleted_by" gorm:"type:varchar(255)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	LastSync      *time.Time     `json:"last_sync"`
}

type ProductVariantRepository interface {
	Create(ctx context.Context, variant *ProductVariant) error
	GetByID(ctx context.Context, id uuid.UUID) (*ProductVariant, error)
	Fetch(ctx context.Context, limit int, offset int) ([]ProductVariant, int64, error)
	FetchByProductID(ctx context.Context, productID int) ([]ProductVariant, error)
	Update(ctx context.Context, variant *ProductVariant) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ProductVariantUseCase interface {
	Create(ctx context.Context, variant *ProductVariant) error
	GetByID(ctx context.Context, id uuid.UUID) (*ProductVariant, error)
	Fetch(ctx context.Context, limit int, offset int) ([]ProductVariant, int64, error)
	FetchByProductID(ctx context.Context, productID int) ([]ProductVariant, error)
	Update(ctx context.Context, variant *ProductVariant) error
	Delete(ctx context.Context, id uuid.UUID) error
}
