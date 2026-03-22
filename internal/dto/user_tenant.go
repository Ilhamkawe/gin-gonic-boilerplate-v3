package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type UserTenantDTO struct {
	UserID   int `json:"user_id" binding:"required"`
	TenantID int `json:"tenant_id" binding:"required"`
	RoleID   int `json:"role_id" binding:"required"`
}

type UserTenantResponse struct {
	ID        int            `json:"id"`
	UUID      uuid.UUID      `json:"uuid"`
	User      *UserResponse  `json:"user,omitempty"`
	Tenant    *TenantPassDTO `json:"tenant,omitempty"`
	Role      *RoleResponse  `json:"role,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func FromUserTenant(ut domain.UserTenant) UserTenantResponse {
	res := UserTenantResponse{
		ID:        ut.ID,
		UUID:      ut.UUID,
		CreatedAt: ut.CreatedAt,
		UpdatedAt: ut.UpdatedAt,
	}

	if ut.User.ID != 0 {
		userRes := FromUser(ut.User)
		res.User = &userRes
	}

	if ut.Tenant.ID != 0 {
		res.Tenant = &TenantPassDTO{
			ID:        ut.Tenant.ID,
			UUID:      ut.Tenant.UUID,
			Name:      ut.Tenant.Name,
			Subdomain: ut.Tenant.Subdomain,
		}
	}

	if ut.Role.ID != 0 {
		res.Role = &RoleResponse{
			ID:   ut.Role.ID,
			UUID: ut.Role.UUID,
			Name: ut.Role.Name,
		}
	}

	return res
}

func FromUserTenants(uts []domain.UserTenant) []UserTenantResponse {
	var responses []UserTenantResponse
	for _, ut := range uts {
		responses = append(responses, FromUserTenant(ut))
	}
	return responses
}
