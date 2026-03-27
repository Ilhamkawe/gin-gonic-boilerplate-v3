package usecase

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type merchantUseCase struct {
	merchantRepo   domain.MerchantRepository
	storageService domain.StorageService
}

func NewMerchantUseCase(merchantRepo domain.MerchantRepository, storageService domain.StorageService) domain.MerchantUseCase {
	return &merchantUseCase{merchantRepo: merchantRepo, storageService: storageService}
}

func (u *merchantUseCase) Create(ctx context.Context, merchant *domain.Merchant, tenantUUID string) error {
	UUID := uuid.New()
	merchant.UUID = UUID

	if merchant.Photo != "" {
		parts := strings.Split(merchant.Photo, "/")
		fileName := parts[len(parts)-1]
		
		sourcePath := ""
		for i, part := range parts {
			if part == "temp" {
				sourcePath = strings.Join(parts[i:], "/")
				break
			}
		}

		if sourcePath != "" {
			destPath := tenantUUID + "/merchant/" + fileName
			err := u.storageService.MoveFile(ctx, sourcePath, destPath)
			if err != nil {
				return err
			}
			merchant.Photo = strings.Replace(merchant.Photo, sourcePath, destPath, 1)
		}
	}

	return u.merchantRepo.Create(ctx, merchant)
}

func (u *merchantUseCase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Merchant, error) {
	return u.merchantRepo.GetByID(ctx, id)
}

func (u *merchantUseCase) Fetch(ctx context.Context, limit int, offset int) ([]domain.Merchant, int64, error) {
	return u.merchantRepo.Fetch(ctx, limit, offset)
}

func (u *merchantUseCase) Update(ctx context.Context, merchant *domain.Merchant, tenantUUID string) error {
	if merchant.Photo != "" && strings.Contains(merchant.Photo, "/temp/") {
		parts := strings.Split(merchant.Photo, "/")
		fileName := parts[len(parts)-1]
		
		sourcePath := ""
		for i, part := range parts {
			if part == "temp" {
				sourcePath = strings.Join(parts[i:], "/")
				break
			}
		}

		if sourcePath != "" {
			destPath := tenantUUID + "/merchant/" + fileName
			err := u.storageService.MoveFile(ctx, sourcePath, destPath)
			if err != nil {
				return err
			}
			merchant.Photo = strings.Replace(merchant.Photo, sourcePath, destPath, 1)
		}
	}

	return u.merchantRepo.Update(ctx, merchant)
}

func (u *merchantUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.merchantRepo.Delete(ctx, id)
}
