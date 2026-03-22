package dto

import "github.com/google/uuid"

type RolePermissionDto struct {
	ID   int       `json:"id"`
	UUID uuid.UUID `json:"uuid"`
}
