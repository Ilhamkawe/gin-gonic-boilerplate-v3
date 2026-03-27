package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type TenantDTO struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID      uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"not null;type:varchar(255)"`
	Address   string    `json:"address" gorm:"not null;type:varchar(255)"`
	Phone     string    `json:"phone" gorm:"not null;type:varchar(15)"`
	Email     string    `json:"email" gorm:"not null;type:varchar(255)"`
	Photo     string    `json:"photo" gorm:"not null;type:varchar(255)"`
	Subdomain string    `json:"subdomain" gorm:"not null;type:varchar(255)"`
	CreatedBy string    `json:"created_by" gorm:"type:varchar(255);not null"`
	UpdatedBy string    `json:"updated_by" gorm:"type:varchar(255)"`
	DeletedBy string    `json:"deleted_by" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	LastSync  time.Time `json:"last_sync"`
}

type TenantResponseDTO struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID      uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"not null;type:varchar(255)"`
	Address   string    `json:"address" gorm:"not null;type:varchar(255)"`
	Phone     string    `json:"phone" gorm:"not null;type:varchar(15)"`
	Email     string    `json:"email" gorm:"not null;type:varchar(255)"`
	Photo     string    `json:"photo" gorm:"not null;type:varchar(255)"`
	Subdomain string    `json:"subdomain" gorm:"not null;type:varchar(255)"`
	OwnerId   int       `json:"owner_id" gorm:"not null"`
}

type TenantPassDTO struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;unique"`
	UUID      uuid.UUID `json:"uuid" gorm:"type:uuid;not null;unique;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"not null;type:varchar(255)"`
	Subdomain string    `json:"subdomain" gorm:"not null;type:varchar(255)"`
}

type CreateTenantDTO struct {
	Name      string `json:"name" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Photo     string `json:"photo"`
	Subdomain string `json:"subdomain" validate:"required"`
}

type UpdateTenantDTO struct {
	ID        int       `json:"id"`
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name" validate:"required"`
	Address   string    `json:"address" validate:"required"`
	Phone     string    `json:"phone" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Photo     string    `json:"photo"`
	Subdomain string    `json:"subdomain" validate:"required"`
}

func FromTenant(tenant domain.Tenant) TenantResponseDTO {
	return TenantResponseDTO{
		ID:        tenant.ID,
		UUID:      tenant.UUID,
		Name:      tenant.Name,
		Address:   tenant.Address,
		Phone:     tenant.Phone,
		Email:     tenant.Email,
		Photo:     tenant.Photo,
		Subdomain: tenant.Subdomain,
		OwnerId:   tenant.OwnerId,
	}
}
