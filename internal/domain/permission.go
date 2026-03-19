package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID        uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique"`
	Name        string    `json:"name" gorm:"not null;type:varchar(255)"`
	Module      string    `json:"module" gorm:"not null;type:varchar(255)"`
	Description string    `json:"description" gorm:"not null;type:varchar(255)"`
	CreatedBy   string    `json:"created_by" gorm:"type:varchar(255);not null"`
	UpdatedBy   string    `json:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy   string    `json:"deleted_by" gorm:"type:varchar(255)"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	LastSync    time.Time `json:"last_sync"`
}

type PermissionRepository interface {
	Create(ctx context.Context, permission *Permission) error
	GetByID(ctx context.Context, id uuid.UUID) (*Permission, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Permission, int64, error)
	Update(ctx context.Context, permission *Permission) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type PermissionUseCase interface {
	Create(ctx context.Context, permission *Permission) error
	GetByID(ctx context.Context, id uuid.UUID) (*Permission, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Permission, int64, error)
	Update(ctx context.Context, permission *Permission) error
	Delete(ctx context.Context, id uuid.UUID) error
}
