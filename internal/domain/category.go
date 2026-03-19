package domain

import (
	"context"
	"io"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Category struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID      uuid.UUID      `json:"uuid" gorm:"type:uuid;not null;unique"`
	Name      string         `json:"name" gorm:"not null;type:varchar(255)"`
	Icon      string         `json:"icon" gorm:"not null;type:varchar(255)"`
	Tagline   string         `json:"tagline" gorm:"not null;type:varchar(255)"`
	FormJson  datatypes.JSON `json:"form_json" gorm:"type:jsonb;not null"`
	TenantID  int            `json:"tenant_id" gorm:"not null"`
	Tenant    Tenant         `json:"tenant" gorm:"foreignKey:TenantID"`
	CreatedBy string         `json:"created_by" gorm:"type:varchar(255);not null"`
	UpdatedBy string         `json:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy string         `json:"deleted_by" gorm:"type:varchar(255)"`
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
	Create(ctx context.Context, category *Category, file io.Reader, fileSize int64) error
	GetByID(ctx context.Context, id uuid.UUID) (*Category, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Category, int64, error)
	Update(ctx context.Context, category *Category, file io.Reader, fileSize int64) error
	Delete(ctx context.Context, id uuid.UUID) error
}
