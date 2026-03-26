package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type rolePermissionUseCase struct {
	rpRepo domain.RolePermissionRepository
}

func NewRolePermissionUseCase(rpRepo domain.RolePermissionRepository) domain.RolePermissionUseCase {
	return &rolePermissionUseCase{rpRepo: rpRepo}
}

func (u *rolePermissionUseCase) Create(ctx context.Context, rp *domain.RolePermission) error {
	rp.UUID = uuid.New()
	return u.rpRepo.Create(ctx, rp)
}

func (u *rolePermissionUseCase) GetByID(ctx context.Context, id uuid.UUID) (*domain.RolePermission, error) {
	return u.rpRepo.GetByID(ctx, id)
}

func (u *rolePermissionUseCase) Fetch(ctx context.Context, limit int, offset int) ([]domain.RolePermission, int64, error) {
	return u.rpRepo.Fetch(ctx, limit, offset)
}

func (u *rolePermissionUseCase) Update(ctx context.Context, rp *domain.RolePermission) error {
	return u.rpRepo.Update(ctx, rp)
}

func (u *rolePermissionUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.rpRepo.Delete(ctx, id)
}

func (u *rolePermissionUseCase) BulkInsert(ctx context.Context, rps []domain.RolePermission) error {
	for i := range rps {
		rps[i].UUID = uuid.New()
	}
	return u.rpRepo.BulkInsert(ctx, rps)
}
