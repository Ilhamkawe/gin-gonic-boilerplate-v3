package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type warehouseUseCase struct {
	warehouseRepo domain.WarehouseRepository
}

func NewWarehouseUseCase(warehouseRepo domain.WarehouseRepository) domain.WarehouseUseCase {
	return &warehouseUseCase{warehouseRepo: warehouseRepo}
}

func (u *warehouseUseCase) Create(ctx context.Context, warehouse *domain.Warehouse) error {
	return u.warehouseRepo.Create(ctx, warehouse)
}

func (u *warehouseUseCase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Warehouse, error) {
	return u.warehouseRepo.GetByID(ctx, id)
}

func (u *warehouseUseCase) Fetch(ctx context.Context, limit int, offset int) ([]domain.Warehouse, int64, error) {
	return u.warehouseRepo.Fetch(ctx, limit, offset)
}

func (u *warehouseUseCase) Update(ctx context.Context, warehouse *domain.Warehouse) error {
	return u.warehouseRepo.Update(ctx, warehouse)
}

func (u *warehouseUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.warehouseRepo.Delete(ctx, id)
}
