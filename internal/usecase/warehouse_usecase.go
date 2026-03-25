package usecase

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type warehouseUseCase struct {
	warehouseRepo  domain.WarehouseRepository
	storageService domain.StorageService
}

func NewWarehouseUseCase(warehouseRepo domain.WarehouseRepository, storageService domain.StorageService) domain.WarehouseUseCase {
	return &warehouseUseCase{warehouseRepo: warehouseRepo, storageService: storageService}
}

func (u *warehouseUseCase) Create(ctx context.Context, warehouse *domain.Warehouse, file io.Reader, fileSize int64) error {
	UUID := uuid.New()
	fileName := UUID.String() + "/warehouses/" + UUID.String() + ".jpg"
	imageUrl, err := u.storageService.UploadFile(ctx, fileName, file, fileSize, "image/jpeg")
	if err != nil {
		return err
	}

	warehouse.UUID = UUID
	warehouse.Photo = imageUrl

	return u.warehouseRepo.Create(ctx, warehouse)
}

func (u *warehouseUseCase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Warehouse, error) {
	return u.warehouseRepo.GetByID(ctx, id)
}

func (u *warehouseUseCase) Fetch(ctx context.Context, limit int, offset int) ([]domain.Warehouse, int64, error) {
	return u.warehouseRepo.Fetch(ctx, limit, offset)
}

func (u *warehouseUseCase) Update(ctx context.Context, warehouse *domain.Warehouse, file io.Reader, fileSize int64) error {
	if file != nil {
		UUID := uuid.New()
		fileName := UUID.String() + "/warehouses/" + UUID.String() + ".jpg"
		imageUrl, err := u.storageService.UploadFile(ctx, fileName, file, fileSize, "image/jpeg")
		if err != nil {
			return err
		}
		warehouse.Photo = imageUrl
	}

	return u.warehouseRepo.Update(ctx, warehouse)
}

func (u *warehouseUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.warehouseRepo.Delete(ctx, id)
}
