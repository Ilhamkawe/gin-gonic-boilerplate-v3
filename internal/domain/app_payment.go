package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type AppPayment struct {
	ID               int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID             uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique"`
	TenantID         int       `json:"tenant_id" gorm:"not null"`
	Tenant           Tenant    `gorm:"foreignKey:TenantID;references:ID"`
	PaymentStatus    string    `json:"payment_status" gorm:"not null; type:varchar(255)"`
	PaymentMethod    string    `json:"payment_method" gorm:"not null; type:varchar(255)"`
	BilingCycleCount int       `json:"biling_cycle_count" gorm:"not null"`
	PlanID           int       `json:"plan_id" gorm:"not null"`
	Plan             Plan      `gorm:"foreignKey:PlanID;references:ID"`
	CreatedBy        string    `json:"created_by" gorm:"type:varchar(255);not null"`
	UpdatedBy        string    `json:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy        string    `json:"deleted_by" gorm:"type:varchar(255)"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        time.Time `json:"deleted_at"`
	LastSync         time.Time `json:"last_sync"`
}

type AppPaymentRepository interface {
	CreateAppPayment(ctx context.Context, appPayment *AppPayment) error
	UpdateAppPayment(ctx context.Context, appPayment *AppPayment) error
	DeleteAppPayment(ctx context.Context, appPayment *AppPayment) error
	GetAppPaymentByID(ctx context.Context, id int) (*AppPayment, error)
	GetAppPaymentByUUID(ctx context.Context, uuid uuid.UUID) (*AppPayment, error)
	GetAppPayments(ctx context.Context) ([]AppPayment, error)
}

type AppPaymentService interface {
	CreateAppPayment(ctx context.Context, appPayment *AppPayment) error
	UpdateAppPayment(ctx context.Context, appPayment *AppPayment) error
	DeleteAppPayment(ctx context.Context, appPayment *AppPayment) error
	GetAppPaymentByID(ctx context.Context, id int) (*AppPayment, error)
	GetAppPaymentByUUID(ctx context.Context, uuid uuid.UUID) (*AppPayment, error)
	GetAppPayments(ctx context.Context) ([]AppPayment, error)
}
