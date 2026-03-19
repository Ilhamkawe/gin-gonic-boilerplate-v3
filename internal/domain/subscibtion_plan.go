package domain

import (
	"time"

	"github.com/google/uuid"
)

type Plan struct {
	ID             int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID           uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique"`
	Name           string    `json:"name" gorm:"not null"`
	BilitCycle     string    `json:"bilit_cycle" gorm:"not null; type:enum('monthly','yearly')"`
	Price          float64   `json:"price" gorm:"not null; type:decimal(10,2)"`
	IsActive       bool      `json:"is_active" gorm:"not null;default:true"`
	MaxWarehouse   int       `json:"max_warehouse" gorm:"not null"`
	MaxUser        int       `json:"max_user" gorm:"not null"`
	MaxMechant     int       `json:"max_mechant" gorm:"not null"`
	MaxProduct     int       `json:"max_product" gorm:"not null"`
	MaxTransaction int       `json:"max_transaction" gorm:"not null"`
	Discount       float64   `json:"discount" gorm:"not null"`
	PriceTotal     float64   `json:"price_total" gorm:"not null; type:decimal(10,2)"`
	CreatedBy      string    `json:"created_by" gorm:""`
	UpdatedBy      string    `json:"updated_by" gorm:""`
	DeletedBy      string    `json:"deleted_by" gorm:""`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
	LastSync       time.Time `json:"last_sync"`
}

type PlanRepository interface {
	CreatePlan(plan *Plan) error
	UpdatePlan(plan *Plan) error
	DeletePlan(plan *Plan) error
	GetPlanByID(id int) (*Plan, error)
	GetPlanByUUID(uuid uuid.UUID) (*Plan, error)
	GetPlans() ([]Plan, error)
}

type PlanService interface {
	CreatePlan(plan *Plan) error
	UpdatePlan(plan *Plan) error
	DeletePlan(plan *Plan) error
	GetPlanByID(id int) (*Plan, error)
	GetPlanByUUID(uuid uuid.UUID) (*Plan, error)
	GetPlans() ([]Plan, error)
}
