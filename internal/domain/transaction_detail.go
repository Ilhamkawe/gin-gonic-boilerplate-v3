package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type TransactionDetail struct {
	ID            int         `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID          uuid.UUID   `json:"uuid" gorm:"type:uuid;not null;unique"`
	TransactionID int         `json:"transaction_id" gorm:"not null"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID;references:ID"`
	ProductID     int         `json:"product_id" gorm:"not null"`
	Product       Product     `gorm:"foreignKey:ProductID;references:ID"`
	Quantity      int         `json:"quantity" gorm:"not null"`
	Price         float64     `json:"price" gorm:"not null"`
	SubTotal      float64     `json:"sub_total" gorm:"not null"`
	CreatedAt     time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time   `json:"updated_at"`
	DeletedAt     time.Time   `json:"deleted_at"`
	LastSync      time.Time   `json:"last_sync"`
}

type TransactionDetailRepository interface {
	Create(ctx context.Context, transactionDetail *TransactionDetail) error
	GetByID(ctx context.Context, id uuid.UUID) (*TransactionDetail, error)
	Fetch(ctx context.Context, limit int, offset int) ([]TransactionDetail, int64, error)
	Update(ctx context.Context, transactionDetail *TransactionDetail) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type TransactionDetailUseCase interface {
	Create(ctx context.Context, transactionDetail *TransactionDetail) error
	GetByID(ctx context.Context, id uuid.UUID) (*TransactionDetail, error)
	Fetch(ctx context.Context, limit int, offset int) ([]TransactionDetail, int64, error)
	Update(ctx context.Context, transactionDetail *TransactionDetail) error
	Delete(ctx context.Context, id uuid.UUID) error
}
