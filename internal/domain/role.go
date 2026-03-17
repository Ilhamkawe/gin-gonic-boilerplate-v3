package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID      uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	LastSync  time.Time `json:"last_sync"`
}

type RoleRepository interface {
	Create(ctx context.Context, role *Role) error
	GetByID(ctx context.Context, id uuid.UUID) (*Role, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Role, int64, error)
	Update(ctx context.Context, role *Role) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type RoleUsecase interface {
	Create(ctx context.Context, role *Role) error
	GetByID(ctx context.Context, id uuid.UUID) (*Role, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Role, int64, error)
	Update(ctx context.Context, role *Role) error
	Delete(ctx context.Context, id uuid.UUID) error
}
