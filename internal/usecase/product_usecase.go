package usecase

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type productUseCase struct {
	productRepo    domain.ProductRepository
	storageService domain.StorageService
}

func NewProductUseCase(productRepo domain.ProductRepository, storageService domain.StorageService) domain.ProductUseCase {
	return &productUseCase{productRepo: productRepo, storageService: storageService}
}

func (u *productUseCase) Create(ctx context.Context, product *domain.Product, tenantUUID string) error {
	UUID := uuid.New()
	product.UUID = UUID

	if product.Thumbnail != "" {
		parts := strings.Split(product.Thumbnail, "/")
		fileName := parts[len(parts)-1]
		
		sourcePath := ""
		for i, part := range parts {
			if part == "temp" {
				sourcePath = strings.Join(parts[i:], "/")
				break
			}
		}

		if sourcePath != "" {
			destPath := tenantUUID + "/product/" + fileName
			err := u.storageService.MoveFile(ctx, sourcePath, destPath)
			if err != nil {
				return err
			}
			product.Thumbnail = strings.Replace(product.Thumbnail, sourcePath, destPath, 1)
		}
	}

	return u.productRepo.Create(ctx, product)
}

func (u *productUseCase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	return u.productRepo.GetByID(ctx, id)
}

func (u *productUseCase) Fetch(ctx context.Context, limit int, offset int) ([]domain.Product, int64, error) {
	return u.productRepo.Fetch(ctx, limit, offset)
}

func (u *productUseCase) Update(ctx context.Context, product *domain.Product, tenantUUID string) error {
	if product.Thumbnail != "" && strings.Contains(product.Thumbnail, "/temp/") {
		parts := strings.Split(product.Thumbnail, "/")
		fileName := parts[len(parts)-1]
		
		sourcePath := ""
		for i, part := range parts {
			if part == "temp" {
				sourcePath = strings.Join(parts[i:], "/")
				break
			}
		}

		if sourcePath != "" {
			destPath := tenantUUID + "/product/" + fileName
			err := u.storageService.MoveFile(ctx, sourcePath, destPath)
			if err != nil {
				return err
			}
			product.Thumbnail = strings.Replace(product.Thumbnail, sourcePath, destPath, 1)
		}
	}

	return u.productRepo.Update(ctx, product)
}

func (u *productUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.productRepo.Delete(ctx, id)
}
