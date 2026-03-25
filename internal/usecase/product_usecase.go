package usecase

import (
	"context"
	"io"

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

func (u *productUseCase) Create(ctx context.Context, product *domain.Product, file io.Reader, fileSize int64) error {
	UUID := uuid.New()
	fileName := UUID.String() + "/products/" + UUID.String() + ".jpg"
	imageUrl, err := u.storageService.UploadFile(ctx, fileName, file, fileSize, "image/jpeg")
	if err != nil {
		return err
	}

	product.UUID = UUID
	product.Thumbnail = imageUrl

	return u.productRepo.Create(ctx, product)
}

func (u *productUseCase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	return u.productRepo.GetByID(ctx, id)
}

func (u *productUseCase) Fetch(ctx context.Context, limit int, offset int) ([]domain.Product, int64, error) {
	return u.productRepo.Fetch(ctx, limit, offset)
}

func (u *productUseCase) Update(ctx context.Context, product *domain.Product, file io.Reader, fileSize int64) error {
	if file != nil {
		UUID := uuid.New()
		fileName := UUID.String() + "/products/" + UUID.String() + ".jpg"
		imageUrl, err := u.storageService.UploadFile(ctx, fileName, file, fileSize, "image/jpeg")
		if err != nil {
			return err
		}
		product.Thumbnail = imageUrl
	}

	return u.productRepo.Update(ctx, product)
}

func (u *productUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.productRepo.Delete(ctx, id)
}
