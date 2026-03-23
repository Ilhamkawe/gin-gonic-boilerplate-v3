package domain

import (
	"time"

	"github.com/google/uuid"
)

type ProductVariant struct {
	ID            int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID          uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique;default:gen_random_uuid()"`
	ProductID     int       `json:"product_id" gorm:"not null"`
	Product       Product   `gorm:"foreignKey:ProductID;references:ID"`
	AttributeJson string    `json:"attribute_json" gorm:"type:jsonb"`
	TenantID      int       `json:"tenant_id" gorm:"not null"`
	Tenant        Tenant    `gorm:"foreignKey:TenantID;references:ID"`
	CreatedBy     string    `json:"created_by" gorm:"type:varchar(255);not null"`
	UpdatedBy     string    `json:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy     string    `json:"deleted_by" gorm:"type:varchar(255)"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
	LastSync      time.Time `json:"last_sync"`
}
