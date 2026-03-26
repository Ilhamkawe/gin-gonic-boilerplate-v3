package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EntityType string

const (
	EntityWarehouse EntityType = "WAREHOUSE"
	EntityOffice    EntityType = "OFFICE"
	EntityStore     EntityType = "STORE"
)

type UserAccess struct {
	ID           int        `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID         uuid.UUID  `json:"uuid" gorm:"type:uuid;not null;unique;default:gen_random_uuid()"`
	UserID       int        `json:"user_id" gorm:"not null"`
	User         User       `gorm:"foreignKey:UserID;references:ID"`
	RoleID       int        `json:"role_id" gorm:"not null"`
	Role         Role       `gorm:"foreignKey:RoleID;references:ID"`
	EntityID     int        `json:"entity_id" gorm:"not null"`
	EntityType   EntityType `json:"entity_type" gorm:"type:string;not null"`
	TenantID     int        `json:"tenant_id" gorm:"not null"`
	Tenant       Tenant     `gorm:"foreignKey:TenantID;references:ID"`
	UserTenantID int        `json:"user_tenant_id" gorm:"not null"`
	UserTenant   UserTenant `gorm:"foreignKey:UserTenantID;references:ID"`
	CreatedBy    string     `json:"created_by" gorm:"type:varchar(255);not null"`
	UpdatedBy    string     `json:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy    string     `json:"deleted_by" gorm:"type:varchar(255)"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	LastSync     time.Time      `json:"last_sync"`
}

type UserAccessRepository interface {
	Create(ctx context.Context, userAccess *UserAccess) error
	GetByID(ctx context.Context, id uuid.UUID) (*UserAccess, error)
	Fetch(ctx context.Context, limit int, offset int) ([]UserAccess, int64, error)
	Update(ctx context.Context, userAccess *UserAccess) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserAccessUseCase interface {
	Create(ctx context.Context, userAccess *UserAccess) error
	GetByID(ctx context.Context, id uuid.UUID) (*UserAccess, error)
	Fetch(ctx context.Context, limit int, offset int) ([]UserAccess, int64, error)
	Update(ctx context.Context, userAccess *UserAccess) error
	Delete(ctx context.Context, id uuid.UUID) error
}
