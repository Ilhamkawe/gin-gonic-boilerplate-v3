package usecase

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type tentantUseCase struct {
	tenantRepo     domain.TenantRepository
	storageService domain.StorageService
}

func NewTenantUseCase(tenantRepo domain.TenantRepository, storageService domain.StorageService) domain.TenantUseCase {
	return &tentantUseCase{tenantRepo: tenantRepo, storageService: storageService}
}

func (t *tentantUseCase) Create(ctx context.Context, tenant *domain.Tenant, file io.Reader, fileSize int64) error {
	UUID := uuid.New()
	fileName := UUID.String() + ".jpg"
	imageUrl, err := t.storageService.UploadFile(ctx, fileName, file, fileSize, "image/jpeg")
	if err != nil {
		return err
	}

	tenant.Photo = imageUrl
	tenant.UUID = UUID

	return t.tenantRepo.Create(ctx, tenant)
}

func (t *tentantUseCase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Tenant, error) {
	return t.tenantRepo.GetByID(ctx, id)
}

func (t *tentantUseCase) Fetch(ctx context.Context, limit int, offset int) ([]domain.Tenant, int64, error) {
	return t.tenantRepo.Fetch(ctx, limit, offset)
}

func (t *tentantUseCase) Update(ctx context.Context, tenant *domain.Tenant, file io.Reader, fileSize int64) error {
	if file != nil {
		UUID := uuid.New()
		fileName := UUID.String() + ".jpg"
		imageUrl, err := t.storageService.UploadFile(ctx, fileName, file, fileSize, "image/jpeg")
		if err != nil {
			return err
		}
		tenant.Photo = imageUrl
	}

	return t.tenantRepo.Update(ctx, tenant)
}

func (t *tentantUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return t.tenantRepo.Delete(ctx, id)
}

func (t *tentantUseCase) IsAuthroized(ctx context.Context, id uuid.UUID, tenantID int) (bool, error) {
	return t.tenantRepo.IsAuthroized(ctx, id, tenantID)
}
