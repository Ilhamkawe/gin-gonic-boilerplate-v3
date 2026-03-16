package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type RolePermission struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UUID         uuid.UUID `json:"uuid" gorm:"type:uuid;primaryKey"`
	RoleID       int
	Role         Role `gorm:"foreignKey:RoleID;references:ID"`
	PermissionID int
	Permission   Permission `gorm:"foreignKey:PermissionID;references:ID"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    time.Time  `json:"deleted_at"`
	LastSync     time.Time  `json:"last_sync"`
}

type RolePermissionRepository interface {
	Create(ctx context.Context, rolePermission *RolePermission) error
	GetByID(ctx context.Context, id uuid.UUID) (*RolePermission, error)
	Fetch(ctx context.Context, limit int, offset int) ([]RolePermission, int64, error)
	Update(ctx context.Context, rolePermission *RolePermission) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type RolePermissionUseCasee interface {
	Create(ctx context.Context, rolePermission *RolePermission) error
	GetByID(ctx context.Context, id uuid.UUID) (*RolePermission, error)
	Fetch(ctx context.Context, limit int, offset int) ([]RolePermission, int64, error)
	Update(ctx context.Context, rolePermission *RolePermission) error
	Delete(ctx context.Context, id uuid.UUID) error
}
