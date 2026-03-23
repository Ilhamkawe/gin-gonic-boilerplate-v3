package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         int          `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID       uuid.UUID    `json:"uuid" gorm:"type:uuid;not null;unique;default:gen_random_uuid()"`
	Email      string       `json:"email" gorm:"unique;not null;type:varchar(255)"`
	Password   string       `json:"-" gorm:"not null;type:varchar(255)"`
	Name       string       `json:"name" gorm:"not null;type:varchar(255)"`
	Photo      string       `json:"photo" gorm:"not null;type:varchar(255)"`
	Phone      string       `json:"phone" gorm:"not null;type:varchar(15)"`
	UserTenant []UserTenant `json:"user_tenant" gorm:"not null;foreignKey:UserID;references:ID"`
	CreatedBy  string       `json:"created_by" gorm:"not null;type:varchar(255)"`
	UpdatedBy  string       `json:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy  string       `json:"deleted_by" gorm:"type:varchar(255)"`
	CreatedAt  time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time    `json:"updated_at"`
	DeletedAt  time.Time    `json:"deleted_at"`
	LastSync   time.Time    `json:"last_sync"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*User, error)
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
	GenerateToken(ctx context.Context, id uuid.UUID, tenantID uuid.UUID) (string, error)
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*User, error)
}
