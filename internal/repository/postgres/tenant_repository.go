package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"gorm.io/gorm"
)

type tenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) domain.TenantRepository {
	return &tenantRepository{db: db}
}

func (t *tenantRepository) Create(ctx context.Context, tenant *domain.Tenant) error {
	return t.db.Create(tenant).Error
}

func (t *tenantRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Tenant, error) {
	var tenant domain.Tenant
	if err := t.db.Where("uuid = ?", id).First(&tenant).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (t *tenantRepository) Fetch(ctx context.Context, userID int, limit int, offset int) ([]domain.Tenant, int64, error) {
	var tenants []domain.Tenant
	var count int64

	query := t.db.Model(&domain.Tenant{}).
		Joins("JOIN user_tenants ON tenants.id = user_tenants.tenant_id").
		Where("user_tenants.user_id = ? AND user_tenants.deleted_at IS NULL", userID)

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Find(&tenants).Error; err != nil {
		return nil, 0, err
	}

	return tenants, count, nil
}

func (t *tenantRepository) Update(ctx context.Context, tenant *domain.Tenant) error {
	return t.db.Model(&domain.Tenant{}).Where("uuid = ?", tenant.UUID).Updates(tenant).Error
}

func (t *tenantRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return t.db.Delete(&domain.Tenant{}, "uuid = ?", id).Error
}

func (t *tenantRepository) IsAuthorized(ctx context.Context, id uuid.UUID, userID int) (bool, error) {
	var count int64
	if err := t.db.Model(&domain.Tenant{}).
		Joins("JOIN user_tenants ON tenants.id = user_tenants.tenant_id").
		Where("tenants.uuid = ? AND user_tenants.user_id = ? AND user_tenants.deleted_at IS NULL", id, userID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (t *tenantRepository) GetAuthorizedTenant(ctx context.Context, tenantID uuid.UUID, userID int) (domain.Tenant, error) {
	var tenant domain.Tenant
	if err := t.db.Joins("JOIN user_tenants ON tenants.id = user_tenants.tenant_id").
		Where("tenants.uuid = ? AND user_tenants.user_id = ? AND user_tenants.deleted_at IS NULL", tenantID, userID).
		First(&tenant).Error; err != nil {
		return tenant, err
	}
	return tenant, nil
}

func (t *tenantRepository) GetAuthorizedTenants(ctx context.Context, userID int) ([]domain.Tenant, error) {
	var tenants []domain.Tenant
	if err := t.db.Joins("JOIN user_tenants ON tenants.id = user_tenants.tenant_id").
		Where("user_tenants.user_id = ? AND user_tenants.deleted_at IS NULL", userID).
		Find(&tenants).Error; err != nil {
		return nil, err
	}
	return tenants, nil
}

func (t *tenantRepository) GetBySubdomain(ctx context.Context, subdomain string) (*domain.Tenant, error) {
	var tenant domain.Tenant
	if err := t.db.Where("subdomain = ?", subdomain).First(&tenant).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (t *tenantRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*domain.Tenant, error) {
	var tenant domain.Tenant
	if err := t.db.Where("uuid = ?", uuid).First(&tenant).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}
