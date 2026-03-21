package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Tenant struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID      uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"not null;type:varchar(255)"`
	Address   string    `json:"address" gorm:"not null;type:varchar(255)"`
	Phone     string    `json:"phone" gorm:"not null;type:varchar(15)"`
	Email     string    `json:"email" gorm:"not null;type:varchar(255)"`
	Photo     string    `json:"photo" gorm:"not null;type:varchar(255)"`
	OwnerId   int       `json:"owner_id" gorm:"not null"`
	Owner     User      `gorm:"foreignKey:OwnerId;references:ID"`
	Subdomain string    `json:"subdomain" gorm:"not null;type:varchar(255)"`
	CreatedBy string    `json:"created_by" gorm:"type:varchar(255);not null"`
	UpdatedBy string    `json:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy string    `json:"deleted_by" gorm:"type:varchar(255)"`
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
	IsAuthroized(ctx context.Context, id uuid.UUID, tenantID int) (bool, error)
}
