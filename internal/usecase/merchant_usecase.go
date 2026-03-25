package usecase

import (
	"context"
	"io"

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

func (u *merchantUseCase) Create(ctx context.Context, merchant *domain.Merchant, file io.Reader, fileSize int64) error {
	UUID := uuid.New()
	fileName := UUID.String() + "/merchants/" + UUID.String() + ".jpg"
	imageUrl, err := u.storageService.UploadFile(ctx, fileName, file, fileSize, "image/jpeg")
	if err != nil {
		return err
	}

	merchant.UUID = UUID
	merchant.Photo = imageUrl

	return u.merchantRepo.Create(ctx, merchant)
}

func (u *merchantUseCase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Merchant, error) {
	return u.merchantRepo.GetByID(ctx, id)
}

func (u *merchantUseCase) Fetch(ctx context.Context, limit int, offset int) ([]domain.Merchant, int64, error) {
	return u.merchantRepo.Fetch(ctx, limit, offset)
}

func (u *merchantUseCase) Update(ctx context.Context, merchant *domain.Merchant, file io.Reader, fileSize int64) error {
	if file != nil {
		UUID := uuid.New()
		fileName := UUID.String() + "/merchants/" + UUID.String() + ".jpg"
		imageUrl, err := u.storageService.UploadFile(ctx, fileName, file, fileSize, "image/jpeg")
		if err != nil {
			return err
		}
		merchant.Photo = imageUrl
	}

	return u.merchantRepo.Update(ctx, merchant)
}

func (u *merchantUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.merchantRepo.Delete(ctx, id)
}
