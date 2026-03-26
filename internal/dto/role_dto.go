package dto

import (
	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type RoleResponse struct {
	ID   int       `json:"id"`
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

type CreateRoleDTO struct {
	Name string `json:"name" validate:"required"`
}

type UpdateRoleDTO struct {
	Name string `json:"name" validate:"required"`
}

func FromRole(role domain.Role) RoleResponse {
	return RoleResponse{
		ID:   role.ID,
		UUID: role.UUID,
		Name: role.Name,
	}
}

func FromRoles(roles []domain.Role) []RoleResponse {
	responses := make([]RoleResponse, 0)
	for _, role := range roles {
		responses = append(responses, FromRole(role))
	}
	return responses
}
