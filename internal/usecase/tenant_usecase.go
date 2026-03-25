package usecase

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type tentantUseCase struct {
	tenantRepo        domain.TenantRepository
	userTenantUseCase domain.UserTenantUseCase
	roleUsecase       domain.RoleUsecase
	storageService    domain.StorageService
}

func NewTenantUseCase(tenantRepo domain.TenantRepository, userTenantUseCase domain.UserTenantUseCase, roleUsecase domain.RoleUsecase, storageService domain.StorageService) domain.TenantUseCase {
	return &tentantUseCase{tenantRepo: tenantRepo, userTenantUseCase: userTenantUseCase, roleUsecase: roleUsecase, storageService: storageService}
}

func (t *tentantUseCase) Create(ctx context.Context, tenant *domain.Tenant, file io.Reader, fileSize int64) error {
	UUID := uuid.New()
	fileName := UUID.String() + "/" + "tenants/" + UUID.String() + ".jpg"
	imageUrl, err := t.storageService.UploadFile(ctx, fileName, file, fileSize, "image/jpeg")
	if err != nil {
		return err
	}

	tenant.Photo = imageUrl
	tenant.UUID = UUID

	err = t.tenantRepo.Create(ctx, tenant)
	if err != nil {
		return err
	}

	defaultRoles := []domain.Role{
		{
			Name:      "owner",
			UUID:      uuid.New(),
			TenantID:  tenant.ID,
			CreatedBy: tenant.CreatedBy,
		},
		{
			Name:      "warehouse_manager",
			UUID:      uuid.New(),
			TenantID:  tenant.ID,
			CreatedBy: tenant.CreatedBy,
		},
		{
			Name:      "merchant_keeper",
			UUID:      uuid.New(),
			TenantID:  tenant.ID,
			CreatedBy: tenant.CreatedBy,
		},
	}

	for i := range defaultRoles {
		if err := t.roleUsecase.Create(ctx, &defaultRoles[i]); err != nil {
			return err
		}
	}

	userTenant := domain.UserTenant{
		UserID:    tenant.OwnerId,
		TenantID:  tenant.ID,
		RoleID:    defaultRoles[0].ID,
		CreatedBy: tenant.CreatedBy,
	}

	err = t.userTenantUseCase.Create(ctx, &userTenant)
	if err != nil {
		return err
	}

	return nil
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
		fileName := UUID.String() + "/" + "tenants/" + UUID.String() + ".jpg"
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

func (t *tentantUseCase) IsAuthorized(ctx context.Context, id uuid.UUID, ownerID int) (bool, error) {
	return t.tenantRepo.IsAuthorized(ctx, id, ownerID)
}

func (t *tentantUseCase) GetAuthorizedTenant(ctx context.Context, tenantID uuid.UUID, ownerID int) (domain.Tenant, error) {
	return t.tenantRepo.GetAuthorizedTenant(ctx, tenantID, ownerID)
}

func (t *tentantUseCase) GetAuthorizedTenants(ctx context.Context, ownerID int) ([]domain.Tenant, error) {
	return t.tenantRepo.GetAuthorizedTenants(ctx, ownerID)
}

func (t *tentantUseCase) GetBySubdomain(ctx context.Context, subdomain string) (*domain.Tenant, error) {
	return t.tenantRepo.GetBySubdomain(ctx, subdomain)
}

func (t *tentantUseCase) IsSubdomainExist(ctx context.Context, subdomain string) (bool, error) {
	_, err := t.tenantRepo.GetBySubdomain(ctx, subdomain)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (t *tentantUseCase) GetByUUID(ctx context.Context, uuid uuid.UUID) (*domain.Tenant, error) {
	return t.tenantRepo.GetByUUID(ctx, uuid)
}
