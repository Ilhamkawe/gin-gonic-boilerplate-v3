package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ProductAttribute struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID       uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique;default:gen_random_uuid()"`
	ProductID  int       `json:"product_id" gorm:"not null"`
	Product    Product   `gorm:"foreignKey:ProductID;references:ID"`
	CategoryID int       `json:"category_id" gorm:"not null"`
	Category   Category  `gorm:"foreignKey:CategoryID;references:ID"`
	Name       string    `json:"name" gorm:"not null;type:varchar(255)"`
	Type       string    `json:"type" gorm:"not null;type:varchar(255)"`
	CreatedBy  string    `json:"created_by" gorm:"type:varchar(255);not null"`
	UpdatedBy  string    `json:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy  string    `json:"deleted_by" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	LastSync   time.Time `json:"last_sync"`
}

type ProductAttributeRepository interface {
	Create(ctx context.Context, productAttribute *ProductAttribute) error
	GetByID(ctx context.Context, id uuid.UUID) (*ProductAttribute, error)
	Fetch(ctx context.Context, limit int, offset int) ([]ProductAttribute, int64, error)
	Update(ctx context.Context, productAttribute *ProductAttribute) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ProductAttributeUseCase interface {
	Create(ctx context.Context, productAttribute *ProductAttribute) error
	GetByID(ctx context.Context, id uuid.UUID) (*ProductAttribute, error)
	Fetch(ctx context.Context, limit int, offset int) ([]ProductAttribute, int64, error)
	Update(ctx context.Context, productAttribute *ProductAttribute) error
	Delete(ctx context.Context, id uuid.UUID) error
}
