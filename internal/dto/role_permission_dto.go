package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type RolePermissionResponse struct {
	ID           int                `json:"id"`
	UUID         uuid.UUID          `json:"uuid"`
	RoleID       int                `json:"role_id"`
	PermissionID int                `json:"permission_id"`
	Permission   PermissionResponse `json:"permission,omitempty"`
	CreatedAt    time.Time          `json:"created_at"`
}

type AssignPermissionsDTO struct {
	RoleID        int   `json:"role_id" validate:"required"`
	PermissionIDs []int `json:"permission_ids" validate:"required"`
}

func FromRolePermission(rp domain.RolePermission) RolePermissionResponse {
	return RolePermissionResponse{
		ID:           rp.ID,
		UUID:         rp.UUID,
		RoleID:       rp.RoleID,
		PermissionID: rp.PermissionID,
		Permission:   FromPermission(rp.Permission),
		CreatedAt:    rp.CreatedAt,
	}
}

func FromRolePermissions(rps []domain.RolePermission) []RolePermissionResponse {
	responses := make([]RolePermissionResponse, 0)
	for _, rp := range rps {
		responses = append(responses, FromRolePermission(rp))
	}
	return responses
}
