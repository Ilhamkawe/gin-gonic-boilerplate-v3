package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type RolePermission struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID         uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique;default:gen_random_uuid()"`
	RoleID       int
	Role         Role `gorm:"foreignKey:RoleID;references:ID"`
	PermissionID int
	Permission   Permission `gorm:"foreignKey:PermissionID;references:ID"`
	TenantID     int        `json:"tenant_id" gorm:"not null"`
	Tenant       Tenant     `gorm:"foreignKey:TenantID;references:ID"`
	CreatedBy    string     `json:"created_by" gorm:"type:varchar(255);not null"`
	UpdatedBy    string     `json:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy    string     `json:"deleted_by" gorm:"type:varchar(255)"`
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
