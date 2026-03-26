package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserTenant struct {
	ID         int          `json:"id" gorm:"primaryKey;autoIncrement"`
	UUID       uuid.UUID    `json:"uuid" gorm:"type:uuid;default:gen_random_uuid()"`
	UserID     int          `json:"user_id" gorm:"not null"`
	User       User         `gorm:"foreignKey:UserID;references:ID"`
	TenantID   int          `json:"tenant_id" gorm:"not null"`
	Tenant     Tenant       `gorm:"foreignKey:TenantID;references:ID"`
	RoleID     int          `json:"role_id" gorm:"not null"`
	Role       Role         `gorm:"foreignKey:RoleID;references:ID"`
	CreatedBy  string       `json:"created_by" gorm:"type:varchar(255);not null"`
	UpdatedBy  string       `json:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy  string       `json:"deleted_by" gorm:"type:varchar(255)"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	LastSync   time.Time      `json:"last_sync"`
	UserAccess []UserAccess `json:"user_access" gorm:"not null;foreignKey:UserTenantID;references:ID"`
}

type UserTenantRepository interface {
	Create(ctx context.Context, userTenant *UserTenant) error
	Update(ctx context.Context, userTenant *UserTenant) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*UserTenant, error)
	GetAll(ctx context.Context) ([]UserTenant, error)
	Fetch(ctx context.Context, limit int, offset int) ([]UserTenant, int64, error)
}

type UserTenantUseCase interface {
	Create(ctx context.Context, userTenant *UserTenant) error
	Update(ctx context.Context, userTenant *UserTenant) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*UserTenant, error)
	GetAll(ctx context.Context) ([]UserTenant, error)
}
