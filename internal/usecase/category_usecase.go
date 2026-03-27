package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type categoryUsecase struct {
	categoryRepo   domain.CategoryRepository
	storageService domain.StorageService
}

func NewCategoryUsecase(categoryRepo domain.CategoryRepository, storageService domain.StorageService) *categoryUsecase {
	return &categoryUsecase{categoryRepo: categoryRepo, storageService: storageService}
}

func (u *categoryUsecase) Create(ctx context.Context, category *domain.Category) error {
	UUID := uuid.New()
	category.UUID = UUID

	if category.Icon != "" {
		parts := strings.Split(category.Icon, "/")
		fileName := parts[len(parts)-1]
		
		sourcePath := ""
		for i, part := range parts {
			if part == "temp" {
				sourcePath = strings.Join(parts[i:], "/")
				break
			}
		}

		if sourcePath != "" {
			destPath := "categories/" + fileName
			err := u.storageService.MoveFile(ctx, sourcePath, destPath)
			if err != nil {
				return err
			}
			category.Icon = strings.Replace(category.Icon, sourcePath, destPath, 1)
		}
	}

	return u.categoryRepo.Create(ctx, category)
}

func (u *categoryUsecase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	return u.categoryRepo.GetByID(ctx, &domain.Category{UUID: id})
}

func (u *categoryUsecase) Fetch(ctx context.Context, limit int, offset int) ([]domain.Category, int64, error) {
	return u.categoryRepo.Fetch(ctx, limit, offset)
}

func (u *categoryUsecase) Update(ctx context.Context, category *domain.Category) error {
	if category.Icon != "" && strings.Contains(category.Icon, "/temp/") {
		parts := strings.Split(category.Icon, "/")
		fileName := parts[len(parts)-1]
		
		sourcePath := ""
		for i, part := range parts {
			if part == "temp" {
				sourcePath = strings.Join(parts[i:], "/")
				break
			}
		}

		if sourcePath != "" {
			destPath := "categories/" + fileName
			err := u.storageService.MoveFile(ctx, sourcePath, destPath)
			if err != nil {
				return err
			}
			category.Icon = strings.Replace(category.Icon, sourcePath, destPath, 1)
		}
	}

	return u.categoryRepo.Update(ctx, category)
}

func (u *categoryUsecase) Delete(ctx context.Context, uuid uuid.UUID) error {

	category, err := u.categoryRepo.GetByID(ctx, &domain.Category{
		UUID:     uuid,
		TenantID: ctx.Value("tenant_id").(int),
	})

	if err != nil {
		return err
	}

	if category.ID == 0 {
		return errors.New("category not found")
	}

	fileName := strings.Split(category.Icon, "/")[len(strings.Split(category.Icon, "/"))-1]

	if err := u.storageService.DeleteFile(ctx, fileName); err != nil {
		return err
	}

	return u.categoryRepo.Delete(ctx, category)
}

func (u *categoryUsecase) GetInsight(ctx context.Context) (*domain.InsightCategory, error) {
	return u.categoryRepo.GetInsight(ctx)
}

func (u *categoryUsecase) IsAvailable(ctx context.Context, uuid uuid.UUID) (bool, error) {
	category, err := u.categoryRepo.GetByID(ctx, &domain.Category{UUID: uuid})

	if err != nil {
		return false, err
	}

	if category.ID == 0 {
		return false, nil
	}

	return true, nil
}

func (u *categoryUsecase) FetchWithProductCount(ctx context.Context, tenantID int) ([]domain.CategoryWithCount, error) {
	return u.categoryRepo.FetchWithProductCount(ctx, tenantID)
}
