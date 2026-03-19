package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Tenant struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID      uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique"`
	Name      string    `json:"name" gorm:"not null"`
	Address   string    `json:"address" gorm:"not null"`
	Phone     string    `json:"phone" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Photo     string    `json:"photo" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	LastSync  time.Time `json:"last_sync"`
}

type TenantRepository interface {
	Create(ctx context.Context, tenant *Tenant) error
	GetByID(ctx context.Context, id uuid.UUID) (*Tenant, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Tenant, int64, error)
	Update(ctx context.Context, tenant *Tenant) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type TenantUseCase interface {
	Create(ctx context.Context, tenant *Tenant) error
	GetByID(ctx context.Context, id uuid.UUID) (*Tenant, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Tenant, int64, error)
	Update(ctx context.Context, tenant *Tenant) error
	Delete(ctx context.Context, id uuid.UUID) error
}
