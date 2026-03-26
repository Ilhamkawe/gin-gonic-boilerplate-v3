package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type UserAccessResponse struct {
	ID           int                `json:"id"`
	UUID         uuid.UUID          `json:"uuid"`
	UserID       int                `json:"user_id"`
	RoleID       int                `json:"role_id"`
	EntityID     int                `json:"entity_id"`
	EntityType   domain.EntityType  `json:"entity_type"`
	TenantID     int                `json:"tenant_id"`
	UserTenantID int                `json:"user_tenant_id"`
	CreatedAt    time.Time          `json:"created_at"`
}

type CreateUserAccessDTO struct {
	UserID       int               `json:"user_id" validate:"required"`
	RoleID       int               `json:"role_id" validate:"required"`
	EntityID     int               `json:"entity_id" validate:"required"`
	EntityType   domain.EntityType `json:"entity_type" validate:"required"`
	UserTenantID int               `json:"user_tenant_id" validate:"required"`
}

func FromUserAccess(ua domain.UserAccess) UserAccessResponse {
	return UserAccessResponse{
		ID:           ua.ID,
		UUID:         ua.UUID,
		UserID:       ua.UserID,
		RoleID:       ua.RoleID,
		EntityID:     ua.EntityID,
		EntityType:   ua.EntityType,
		TenantID:     ua.TenantID,
		UserTenantID: ua.UserTenantID,
		CreatedAt:    ua.CreatedAt,
	}
}

func FromUserAccesses(ua []domain.UserAccess) []UserAccessResponse {
	responses := make([]UserAccessResponse, 0)
	for _, a := range ua {
		responses = append(responses, FromUserAccess(a))
	}
	return responses
}
