package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type productVariantUseCase struct {
	variantRepo domain.ProductVariantRepository
}

func NewProductVariantUseCase(variantRepo domain.ProductVariantRepository) domain.ProductVariantUseCase {
	return &productVariantUseCase{variantRepo: variantRepo}
}

func (u *productVariantUseCase) Create(ctx context.Context, variant *domain.ProductVariant) error {
	variant.UUID = uuid.New()
	return u.variantRepo.Create(ctx, variant)
}

func (u *productVariantUseCase) GetByID(ctx context.Context, id uuid.UUID) (*domain.ProductVariant, error) {
	return u.variantRepo.GetByID(ctx, id)
}

func (u *productVariantUseCase) Fetch(ctx context.Context, limit int, offset int) ([]domain.ProductVariant, int64, error) {
	return u.variantRepo.Fetch(ctx, limit, offset)
}

func (u *productVariantUseCase) FetchByProductID(ctx context.Context, productID int) ([]domain.ProductVariant, error) {
	return u.variantRepo.FetchByProductID(ctx, productID)
}

func (u *productVariantUseCase) Update(ctx context.Context, variant *domain.ProductVariant) error {
	return u.variantRepo.Update(ctx, variant)
}

func (u *productVariantUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.variantRepo.Delete(ctx, id)
}
