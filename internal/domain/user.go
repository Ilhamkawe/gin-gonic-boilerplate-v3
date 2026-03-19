package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         int          `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID       uuid.UUID    `json:"uuid" gorm:"type:uuid;not null;unique"`
	Email      string       `json:"email" gorm:"unique;not null"`
	Password   string       `json:"-" gorm:"not null"`
	Name       string       `json:"name" gorm:"not null"`
	Photo      string       `json:"photo" gorm:"not null"`
	Phone      string       `json:"phone" gorm:"not null"`
	UserAccess []UserAccess `json:"user_access" gorm:"not null;foreignKey:UserID;references:ID"`
	TenantID   int          `json:"tenant_id" gorm:"not null"`
	Tenant     Tenant       `gorm:"foreignKey:TenantID;references:ID"`
	CreatedBy  string       `json:"created_by" gorm:"not null"`
	UpdatedBy  string       `json:"updated_by" gorm:""`
	DeletedBy  string       `json:"deleted_by" gorm:""`
	CreatedAt  time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time    `json:"updated_at"`
	DeletedAt  time.Time    `json:"deleted_at"`
	LastSync   time.Time    `json:"last_sync" gorm:"autoUpdateTime"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Fetch(ctx context.Context, limit int, offset int) ([]User, int64, error)
	GetDetailByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserUsecase interface {
	Login(ctx context.Context, email string, password string) (*User, error)
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	Fetch(ctx context.Context, page int, limit int) ([]User, int64, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id uuid.UUID) error
	GenerateToken(id uuid.UUID) (string, error)
}
