package usecase

import (
	"context"
	"strings"

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

func (u *warehouseUseCase) Create(ctx context.Context, warehouse *domain.Warehouse, tenantUUID string) error {
	UUID := uuid.New()
	warehouse.UUID = UUID

	if warehouse.Photo != "" {
		parts := strings.Split(warehouse.Photo, "/")
		fileName := parts[len(parts)-1]
		
		sourcePath := ""
		for i, part := range parts {
			if part == "temp" {
				sourcePath = strings.Join(parts[i:], "/")
				break
			}
		}

		if sourcePath != "" {
			destPath := tenantUUID + "/warehouse/" + fileName
			err := u.storageService.MoveFile(ctx, sourcePath, destPath)
			if err != nil {
				return err
			}
			warehouse.Photo = strings.Replace(warehouse.Photo, sourcePath, destPath, 1)
		}
	}

	return u.warehouseRepo.Create(ctx, warehouse)
}

func (u *warehouseUseCase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Warehouse, error) {
	return u.warehouseRepo.GetByID(ctx, id)
}

func (u *warehouseUseCase) Fetch(ctx context.Context, limit int, offset int) ([]domain.Warehouse, int64, error) {
	return u.warehouseRepo.Fetch(ctx, limit, offset)
}

func (u *warehouseUseCase) Update(ctx context.Context, warehouse *domain.Warehouse, tenantUUID string) error {
	if warehouse.Photo != "" && strings.Contains(warehouse.Photo, "/temp/") {
		parts := strings.Split(warehouse.Photo, "/")
		fileName := parts[len(parts)-1]
		
		sourcePath := ""
		for i, part := range parts {
			if part == "temp" {
				sourcePath = strings.Join(parts[i:], "/")
				break
			}
		}

		if sourcePath != "" {
			destPath := tenantUUID + "/warehouse/" + fileName
			err := u.storageService.MoveFile(ctx, sourcePath, destPath)
			if err != nil {
				return err
			}
			warehouse.Photo = strings.Replace(warehouse.Photo, sourcePath, destPath, 1)
		}
	}

	return u.warehouseRepo.Update(ctx, warehouse)
}

func (u *warehouseUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.warehouseRepo.Delete(ctx, id)
}
