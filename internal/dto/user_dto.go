package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
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
	Phone    string `json:"phone" validate:"required"`
	Photo    string `json:"photo"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" validate:"omitempty"`
	Password string `json:"password" validate:"omitempty,min=6"`
}

func FromUser(user domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		UUID:      user.UUID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func FromUsers(users []domain.User) []UserResponse {
	userResponses := make([]UserResponse, 0)
	for _, user := range users {
		userResponses = append(userResponses, FromUser(user))
	}
	return userResponses
}

type UserProfileResponse struct {
	ID        int                  `json:"id"`
	UUID      uuid.UUID            `json:"uuid"`
	Email     string               `json:"email"`
	Name      string               `json:"name"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	Tenants   []UserTenantResponse `json:"tenants,omitempty"`
}

func FromUserProfile(user domain.User) UserProfileResponse {
	res := UserProfileResponse{
		ID:        user.ID,
		UUID:      user.UUID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	if len(user.UserTenant) > 0 {
		res.Tenants = FromUserTenants(user.UserTenant)
	}

	return res
}
