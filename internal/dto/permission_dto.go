package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type PermissionResponse struct {
	ID          int       `json:"id"`
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Module      string    `json:"module"`
	Description string    `json:"description"`
	IsAddon     bool      `json:"is_addon"`
	AddonID     int       `json:"addon_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreatePermissionDTO struct {
	Name        string `json:"name" validate:"required"`
	Module      string `json:"module" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsAddon     bool   `json:"is_addon"`
	AddonID     int    `json:"addon_id"`
}

type UpdatePermissionDTO struct {
	Name        string `json:"name" validate:"required"`
	Module      string `json:"module" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsAddon     bool   `json:"is_addon"`
	AddonID     int    `json:"addon_id"`
}

func FromPermission(p domain.Permission) PermissionResponse {
	return PermissionResponse{
		ID:          p.ID,
		UUID:        p.UUID,
		Name:        p.Name,
		Module:      p.Module,
		Description: p.Description,
		IsAddon:     p.IsAddon,
		AddonID:     p.AddonID,
		CreatedAt:   p.CreatedAt,
	}
}

func FromPermissions(permissions []domain.Permission) []PermissionResponse {
	responses := make([]PermissionResponse, 0)
	for _, p := range permissions {
		responses = append(responses, FromPermission(p))
	}
	return responses
}
