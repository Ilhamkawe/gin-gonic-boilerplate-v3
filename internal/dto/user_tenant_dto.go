package dto

type CreateUserTenantDTO struct {
	UserID   int `json:"user_id" form:"user_id" validate:"required"`
	TenantID int `json:"tenant_id" form:"tenant_id" validate:"required"`
	RoleID   int `json:"role_id" form:"role_id" validate:"required"`
}

type UpdateUserTenantDTO struct {
	ID       int `json:"id" form:"id"`
	UserID   int `json:"user_id" form:"user_id" validate:"required"`
	TenantID int `json:"tenant_id" form:"tenant_id" validate:"required"`
	RoleID   int `json:"role_id" form:"role_id" validate:"required"`
}

type GetUserTenantDTO struct {
	ID       int `json:"id" form:"id"`
	UserID   int `json:"user_id" form:"user_id" validate:"required"`
	TenantID int `json:"tenant_id" form:"tenant_id" validate:"required"`
	RoleID   int `json:"role_id" form:"role_id" validate:"required"`
}

type DeleteUserTenantDTO struct {
	ID int `json:"id" form:"id"`
}
