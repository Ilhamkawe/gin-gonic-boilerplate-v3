package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"gorm.io/gorm"
)

type merchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) domain.MerchantRepository {
	return &merchantRepository{db: db}
}

func (r *merchantRepository) Create(ctx context.Context, merchant *domain.Merchant) error {
	return r.db.Create(merchant).Error
}

func (r *merchantRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Merchant, error) {
	var merchant domain.Merchant
	err := r.db.Where("uuid = ? AND tenant_id = ?", id, ctx.Value("tenant_id").(int)).First(&merchant).Error
	return &merchant, err
}

func (r *merchantRepository) Fetch(ctx context.Context, limit int, offset int) ([]domain.Merchant, int64, error) {
	var merchants []domain.Merchant
	var count int64
	tenantID := ctx.Value("tenant_id").(int)

	if err := r.db.Model(&domain.Merchant{}).Where("tenant_id = ? AND deleted_at = '0001-01-01'", tenantID).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Where("tenant_id = ? AND deleted_at = '0001-01-01'", tenantID).Limit(limit).Offset(offset).Find(&merchants).Error
	return merchants, count, err
}

func (r *merchantRepository) Update(ctx context.Context, merchant *domain.Merchant) error {
	return r.db.Where("uuid = ? AND tenant_id = ?", merchant.UUID, merchant.TenantID).Updates(merchant).Error
}

func (r *merchantRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Where("uuid = ? AND tenant_id = ?", id, ctx.Value("tenant_id").(int)).Delete(&domain.Merchant{}).Error
}
