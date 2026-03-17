package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID       uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique"`
	SubTotal   float64   `json:"sub_total" gorm:"not null"`
	Tax        float64   `json:"tax" gorm:"not null"`
	Total      float64   `json:"total" gorm:"not null"`
	MerchantID int       `json:"merchant_id" gorm:"not null"`
	Merchant   Merchant  `gorm:"foreignKey:MerchantID;references:ID"`
	Status     string    `json:"status" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	LastSync   time.Time `json:"last_sync"`
}

type TransactionRepository interface {
	Create(ctx context.Context, transaction *Transaction) error
	GetByID(ctx context.Context, id uuid.UUID) (*Transaction, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Transaction, int64, error)
	Update(ctx context.Context, transaction *Transaction) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type TransactionUseCase interface {
	Create(ctx context.Context, transaction *Transaction) error
	GetByID(ctx context.Context, id uuid.UUID) (*Transaction, error)
	Fetch(ctx context.Context, limit int, offset int) ([]Transaction, int64, error)
	Update(ctx context.Context, transaction *Transaction) error
	Delete(ctx context.Context, id uuid.UUID) error
}
