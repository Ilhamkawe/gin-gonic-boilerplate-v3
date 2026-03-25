package domain

import (
	"context"
	"io"
	"time"

	"github.com/google/uuid"
)

type Warehouse struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID      uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"not null;type:varchar(255)"`
	Address   string    `json:"address" gorm:"not null;type:varchar(255)"`
	Phone     string    `json:"phone" gorm:"not null;type:varchar(15)"`
	Email     string    `json:"email" gorm:"not null;type:varchar(255)"`
	Photo     string    `json:"photo" gorm:"not null;type:varchar(255)"`
	TenantID  int       `json:"tenant_id" gorm:"not null"`
	Tenant    Tenant    `gorm:"foreignKey:TenantID;references:ID"`
	CreatedBy string    `json:"created_by" gorm:"type:varchar(255);not null"`
	UpdatedBy string    `json:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy string    `json:"deleted_by" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	LastSync  time.Time `json:"last_sync"`
}

type WarehouseRepository interface {
	Create(ctx context.Context, warehouse *Warehouse) error
	GetByID(ctx context.Context, id uuid.UUID) (*Warehouse, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Warehouse, int64, error)
	Update(ctx context.Context, warehouse *Warehouse) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type WarehouseUseCase interface {
	Create(ctx context.Context, warehouse *Warehouse, file io.Reader, fileSize int64) error
	GetByID(ctx context.Context, id uuid.UUID) (*Warehouse, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Warehouse, int64, error)
	Update(ctx context.Context, warehouse *Warehouse, file io.Reader, fileSize int64) error
	Delete(ctx context.Context, id uuid.UUID) error
}
