package domain

import (
	"time"

	"github.com/google/uuid"
)

type EntityType string

const (
	EntityWarehouse EntityType = "WAREHOUSE"
	EntityOffice    EntityType = "OFFICE"
	EntityStore     EntityType = "STORE"
)

type UserAccess struct {
	ID         int        `json:"id" gorm:"primaryKey;autoIncrement"`
	UUID       uuid.UUID  `json:"uuid" gorm:"type:uuid;unique;not null"`
	UserID     int        `json:"user_id" gorm:"not null"`
	User       User       `gorm:"foreignKey:UserID;references:ID"`
	RoleID     int        `json:"role_id" gorm:"not null"`
	Role       Role       `gorm:"foreignKey:RoleID;references:ID"`
	EntityID   int        `json:"entity_id" gorm:"not null"`
	EntityType EntityType `json:"entity_type" gorm:"type:string;not null"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  time.Time  `json:"deleted_at"`
	LastSync   time.Time  `json:"last_sync"`
}
