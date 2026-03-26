package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type userAccessUseCase struct {
	uaRepo domain.UserAccessRepository
}

func NewUserAccessUseCase(uaRepo domain.UserAccessRepository) domain.UserAccessUseCase {
	return &userAccessUseCase{uaRepo: uaRepo}
}

func (u *userAccessUseCase) Create(ctx context.Context, ua *domain.UserAccess) error {
	ua.UUID = uuid.New()
	return u.uaRepo.Create(ctx, ua)
}

func (u *userAccessUseCase) GetByID(ctx context.Context, id uuid.UUID) (*domain.UserAccess, error) {
	return u.uaRepo.GetByID(ctx, id)
}

func (u *userAccessUseCase) Fetch(ctx context.Context, limit int, offset int) ([]domain.UserAccess, int64, error) {
	return u.uaRepo.Fetch(ctx, limit, offset)
}

func (u *userAccessUseCase) Update(ctx context.Context, ua *domain.UserAccess) error {
	return u.uaRepo.Update(ctx, ua)
}

func (u *userAccessUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.uaRepo.Delete(ctx, id)
}
