package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Category struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID      uuid.UUID      `json:"uuid" gorm:"type:uuid;not null;unique"`
	Name      string         `json:"name" gorm:"not null"`
	Icon      string         `json:"icon" gorm:"not null"`
	Tagline   string         `json:"tagline" gorm:"not null"`
	FormJson  datatypes.JSON `json:"form_json" gorm:"type:jsonb;not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt time.Time      `json:"deleted_at"`
	LastSync  time.Time      `json:"last_sync"`
}

type CategoryRepository interface {
	Create(ctx context.Context, category *Category) error
	GetByID(ctx context.Context, id uuid.UUID) (*Category, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Category, int64, error)
	Update(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type CategoryUseCase interface {
	Create(ctx context.Context, category *Category) error
	GetByID(ctx context.Context, id uuid.UUID) (*Category, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Category, int64, error)
	Update(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id uuid.UUID) error
}
