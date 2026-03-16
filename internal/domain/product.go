package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UUID        uuid.UUID `json:"uuid" gorm:"type:uuid;primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Thumbnail   string    `json:"icon" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Price       int       `json:"price" gorm:"not null"`
	CategoryID  int       `json:"category_id" gorm:"not null"`
	Category    Category  `gorm:"foreignKey:CategoryID;references:ID"`
	IsPopular   bool      `json:"is_popular" gorm:"not null"`
	FormJson    string    `json:"form_json" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	LastSync    time.Time `json:"last_sync"`
}

type ProductRepository interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id uuid.UUID) (*Product, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Product, int64, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ProductUseCase interface {
	Create(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id uuid.UUID) (*Product, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Product, int64, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id uuid.UUID) error
}
