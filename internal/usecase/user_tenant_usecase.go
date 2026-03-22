package usecase

import (
	"context"

	"github.com/kawe/warehouse_backend/internal/domain"
)

type userTenantUseCase struct {
	userTenantRepo domain.UserTenantRepository
}

func NewUserTenantUseCase(userTenantRepo domain.UserTenantRepository) domain.UserTenantUseCase {
	return &userTenantUseCase{userTenantRepo: userTenantRepo}
}

func (t *userTenantUseCase) Create(ctx context.Context, userTenant *domain.UserTenant) error {
	return t.userTenantRepo.Create(ctx, userTenant)
}

func (t *userTenantUseCase) GetByID(ctx context.Context, id int) (*domain.UserTenant, error) {
	return t.userTenantRepo.GetByID(ctx, id)
}

func (t *userTenantUseCase) Fetch(ctx context.Context, limit int, offset int) ([]domain.UserTenant, int64, error) {
	return t.userTenantRepo.Fetch(ctx, limit, offset)
}

func (t *userTenantUseCase) Update(ctx context.Context, userTenant *domain.UserTenant) error {
	return t.userTenantRepo.Update(ctx, userTenant)
}

func (t *userTenantUseCase) Delete(ctx context.Context, id int) error {
	return t.userTenantRepo.Delete(ctx, id)
}

func (t *userTenantUseCase) GetAll(ctx context.Context) ([]domain.UserTenant, error) {
	return t.userTenantRepo.GetAll(ctx)
}
