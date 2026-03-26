package domain

import (
	"context"
	"io"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID            int        `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID          uuid.UUID  `json:"uuid" gorm:"type:uuid;not null;unique;default:gen_random_uuid()"`
	Name          string     `json:"name" gorm:"not null;type:varchar(255)"`
	Thumbnail     string     `json:"icon" gorm:"not null"`
	Description   string     `json:"description" gorm:"not null"`
	Price         int        `json:"price" gorm:"not null"`
	CategoryID    int        `json:"category_id" gorm:"not null"`
	Category      Category   `gorm:"foreignKey:CategoryID;references:ID"`
	IsPopular     bool       `json:"is_popular" gorm:"not null"`
	TenantID      int        `json:"tenant_id" gorm:"not null"`
	Tenant        Tenant     `gorm:"foreignKey:TenantID;references:ID"`
	AttributeJson string     `json:"attribute_json" gorm:"type:jsonb"`
	CreatedBy     string     `json:"created_by" gorm:"type:varchar(255);not null"`
	UpdatedBy     string     `json:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy     string     `json:"deleted_by" gorm:"type:varchar(255)"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	LastSync      *time.Time     `json:"last_sync"`
}

type ProductRepository interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id uuid.UUID) (*Product, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Product, int64, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ProductUseCase interface {
	Create(ctx context.Context, product *Product, file io.Reader, fileSize int64) error
	GetByID(ctx context.Context, id uuid.UUID) (*Product, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Product, int64, error)
	Update(ctx context.Context, product *Product, file io.Reader, fileSize int64) error
	Delete(ctx context.Context, id uuid.UUID) error
}
