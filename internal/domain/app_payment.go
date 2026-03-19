package domain

import (
	"time"

	"github.com/google/uuid"
)

type AppPayment struct {
	ID               int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID             uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique"`
	TenantID         int       `json:"tenant_id" gorm:"not null"`
	Tenant           Tenant    `gorm:"foreignKey:TenantID;references:ID"`
	PaymentStatus    string    `json:"payment_status" gorm:"not null; type:enum('pending','success','failed')"`
	PaymentMethod    string    `json:"payment_method" gorm:"not null; type:enum('manual','auto')"`
	BilingCycleCount int       `json:"biling_cycle_count" gorm:"not null"`
	PlanID           int       `json:"plan_id" gorm:"not null"`
	Plan             Plan      `gorm:"foreignKey:PlanID;references:ID"`
	CreatedBy        string    `json:"created_by" gorm:""`
	UpdatedBy        string    `json:"updated_by" gorm:""`
	DeletedBy        string    `json:"deleted_by" gorm:""`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        time.Time `json:"deleted_at"`
	LastSync         time.Time `json:"last_sync"`
}

type AppPaymentRepository interface {
	CreateAppPayment(appPayment *AppPayment) error
	UpdateAppPayment(appPayment *AppPayment) error
	DeleteAppPayment(appPayment *AppPayment) error
	GetAppPaymentByID(id int) (*AppPayment, error)
	GetAppPaymentByUUID(uuid uuid.UUID) (*AppPayment, error)
	GetAppPayments() ([]AppPayment, error)
}

type AppPaymentService interface {
	CreateAppPayment(appPayment *AppPayment) error
	UpdateAppPayment(appPayment *AppPayment) error
	DeleteAppPayment(appPayment *AppPayment) error
	GetAppPaymentByID(id int) (*AppPayment, error)
	GetAppPaymentByUUID(uuid uuid.UUID) (*AppPayment, error)
	GetAppPayments() ([]AppPayment, error)
}
