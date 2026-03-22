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

func FromRole(role domain.Role) RoleResponse {
	return RoleResponse{
		ID:   role.ID,
		UUID: role.UUID,
		Name: role.Name,
	}
}

func FromRoles(roles []domain.Role) []RoleResponse {
	var responses []RoleResponse
	for _, role := range roles {
		responses = append(responses, FromRole(role))
	}
	return responses
}
