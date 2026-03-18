package usecase

import (
	"context"
	"io"

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

func (u *categoryUsecase) Create(ctx context.Context, category *domain.Category, file io.Reader, fileSize int64) error {
	UUID := uuid.New()
	fileName := UUID.String() + ".jpg"
	imageUrl, err := u.storageService.UploadFile(ctx, fileName, file, fileSize, "image/jpeg")
	if err != nil {
		return err
	}

	category.UUID = UUID
	category.Icon = imageUrl

	return u.categoryRepo.Create(ctx, category)
}

func (u *categoryUsecase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	return u.categoryRepo.GetByID(ctx, id)
}

func (u *categoryUsecase) Fetch(ctx context.Context, limit int, offset int) ([]domain.Category, int64, error) {
	return u.categoryRepo.Fetch(ctx, limit, offset)
}

func (u *categoryUsecase) Update(ctx context.Context, category *domain.Category) error {
	return u.categoryRepo.Update(ctx, category)
}

func (u *categoryUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.categoryRepo.Delete(ctx, id)
}
