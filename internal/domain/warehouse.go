package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Warehouse struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID      uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique"`
	Name      string    `json:"name" gorm:"not null"`
	Address   string    `json:"address" gorm:"not null"`
	Phone     string    `json:"phone" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Photo     string    `json:"photo" gorm:"not null"`
	ManagerId int       `json:"manager_id" gorm:"not null"`
	Manager   User      `gorm:"foreignKey:ManagerId;references:ID"`
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
	Create(ctx context.Context, warehouse *Warehouse) error
	GetByID(ctx context.Context, id uuid.UUID) (*Warehouse, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Warehouse, int64, error)
	Update(ctx context.Context, warehouse *Warehouse) error
	Delete(ctx context.Context, id uuid.UUID) error
}
