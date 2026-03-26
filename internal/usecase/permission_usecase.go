package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type permissionUseCase struct {
	permissionRepo domain.PermissionRepository
}

func NewPermissionUseCase(permissionRepo domain.PermissionRepository) domain.PermissionUseCase {
	return &permissionUseCase{permissionRepo: permissionRepo}
}

func (u *permissionUseCase) Create(ctx context.Context, p *domain.Permission) error {
	p.UUID = uuid.New()
	return u.permissionRepo.Create(ctx, p)
}

func (u *permissionUseCase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Permission, error) {
	return u.permissionRepo.GetByID(ctx, id)
}

func (u *permissionUseCase) Fetch(ctx context.Context, limit int, offset int) ([]domain.Permission, int64, error) {
	return u.permissionRepo.Fetch(ctx, limit, offset)
}

func (u *permissionUseCase) Update(ctx context.Context, p *domain.Permission) error {
	return u.permissionRepo.Update(ctx, p)
}

func (u *permissionUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.permissionRepo.Delete(ctx, id)
}
