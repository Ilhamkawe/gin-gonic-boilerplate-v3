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

func (t *tenantRepository) Fetch(ctx context.Context, limit int, offset int) ([]domain.Tenant, int64, error) {
	var tenants []domain.Tenant
	var count int64
	if err := t.db.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := t.db.Limit(limit).Offset(offset).Find(&tenants).Error; err != nil {
		return nil, 0, err
	}
	return tenants, count, nil
}

func (t *tenantRepository) Update(ctx context.Context, tenant *domain.Tenant) error {
	return t.db.Save(tenant).Error
}

func (t *tenantRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return t.db.Delete(&domain.Tenant{}, "uuid = ?", id).Error
}

func (t *tenantRepository) IsAuthroized(ctx context.Context, id uuid.UUID, tenantID int) (bool, error) {
	var tenant domain.Tenant
	if err := t.db.Where("uuid = ? AND tenant_id = ?", id, tenantID).First(&tenant).Error; err != nil {
		return false, err
	}
	return true, nil
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
