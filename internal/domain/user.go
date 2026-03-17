package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID      uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"`
	Name      string    `json:"name" gorm:"not null"`
	Photo     string    `json:"photo" gorm:"not null"`
	Phone     string    `json:"phone" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	LastSync  time.Time `json:"last_sync" gorm:"autoUpdateTime"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Fetch(ctx context.Context, limit int, offset int) ([]User, int64, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserUsecase interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	Fetch(ctx context.Context, page int, limit int) ([]User, int64, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id uuid.UUID) error
}
