package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        int       `json:"id"`
	UUID      uuid.UUID `json:"uuid"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" validate:"omitempty"`
	Password string `json:"password" validate:"omitempty,min=6"`
}
