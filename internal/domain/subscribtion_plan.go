package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Plan struct {
	ID             int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID           uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique"`
	Name           string    `json:"name" gorm:"not null; type:varchar(255)"`
	BilitCycle     string    `json:"bilit_cycle" gorm:"not null; type:varchar(20)"`
	Price          float64   `json:"price" gorm:"not null; type:decimal(10,2)"`
	IsActive       bool      `json:"is_active" gorm:"not null;default:true"`
	MaxWarehouse   int       `json:"max_warehouse" gorm:"not null"`
	MaxUser        int       `json:"max_user" gorm:"not null"`
	MaxMechant     int       `json:"max_mechant" gorm:"not null"`
	MaxProduct     int       `json:"max_product" gorm:"not null"`
	MaxTransaction int       `json:"max_transaction" gorm:"not null"`
	Discount       float64   `json:"discount" gorm:"not null"`
	PriceTotal     float64   `json:"price_total" gorm:"not null; type:decimal(10,2)"`
	CreatedBy      string    `json:"created_by" gorm:"type:varchar(255);not null"`
	UpdatedBy      string    `json:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy      string    `json:"deleted_by" gorm:"type:varchar(255)"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
	LastSync       time.Time `json:"last_sync"`
}

type PlanRepository interface {
	CreatePlan(ctx context.Context, plan *Plan) error
	UpdatePlan(ctx context.Context, plan *Plan) error
	DeletePlan(ctx context.Context, plan *Plan) error
	GetPlanByID(ctx context.Context, id int) (*Plan, error)
	GetPlanByUUID(ctx context.Context, uuid uuid.UUID) (*Plan, error)
	GetPlans(ctx context.Context) ([]Plan, error)
}

type PlanService interface {
	CreatePlan(ctx context.Context, plan *Plan) error
	UpdatePlan(ctx context.Context, plan *Plan) error
	DeletePlan(ctx context.Context, plan *Plan) error
	GetPlanByID(ctx context.Context, id int) (*Plan, error)
	GetPlanByUUID(ctx context.Context, uuid uuid.UUID) (*Plan, error)
	GetPlans(ctx context.Context) ([]Plan, error)
}
