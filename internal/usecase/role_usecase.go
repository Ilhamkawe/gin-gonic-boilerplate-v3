package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type roleUsecase struct {
	roleRepo domain.RoleRepository
}

func NewRoleUsecase(roleRepo domain.RoleRepository) domain.RoleUsecase {
	return &roleUsecase{roleRepo: roleRepo}
}

func (t *roleUsecase) Create(ctx context.Context, role *domain.Role) error {
	return t.roleRepo.Create(ctx, role)
}

func (t *roleUsecase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Role, error) {
	return t.roleRepo.GetByID(ctx, id)
}

func (t *roleUsecase) Fetch(ctx context.Context, limit int, offset int) ([]domain.Role, int64, error) {
	return t.roleRepo.Fetch(ctx, limit, offset)
}

func (t *roleUsecase) Update(ctx context.Context, role *domain.Role) error {
	return t.roleRepo.Update(ctx, role)
}

func (t *roleUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return t.roleRepo.Delete(ctx, id)
}
