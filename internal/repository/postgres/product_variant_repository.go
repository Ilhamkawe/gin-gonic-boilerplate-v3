package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"gorm.io/gorm"
)

type productVariantRepository struct {
	db *gorm.DB
}

func NewProductVariantRepository(db *gorm.DB) domain.ProductVariantRepository {
	return &productVariantRepository{db: db}
}

func (r *productVariantRepository) Create(ctx context.Context, variant *domain.ProductVariant) error {
	return r.db.Create(variant).Error
}

func (r *productVariantRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.ProductVariant, error) {
	var variant domain.ProductVariant
	err := r.db.Preload("Product").Where("uuid = ? AND tenant_id = ?", id, ctx.Value("tenant_id").(int)).First(&variant).Error
	return &variant, err
}

func (r *productVariantRepository) Fetch(ctx context.Context, limit int, offset int) ([]domain.ProductVariant, int64, error) {
	var variants []domain.ProductVariant
	var count int64
	tenantID := ctx.Value("tenant_id").(int)

	if err := r.db.Model(&domain.ProductVariant{}).Where("tenant_id = ?", tenantID).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Preload("Product").Where("tenant_id = ?", tenantID).Limit(limit).Offset(offset).Find(&variants).Error
	return variants, count, err
}

func (r *productVariantRepository) FetchByProductID(ctx context.Context, productID int) ([]domain.ProductVariant, error) {
	var variants []domain.ProductVariant
	tenantID := ctx.Value("tenant_id").(int)
	err := r.db.Where("product_id = ? AND tenant_id = ?", productID, tenantID).Find(&variants).Error
	return variants, err
}

func (r *productVariantRepository) Update(ctx context.Context, variant *domain.ProductVariant) error {
	return r.db.Where("uuid = ? AND tenant_id = ?", variant.UUID, variant.TenantID).Updates(variant).Error
}

func (r *productVariantRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Where("uuid = ? AND tenant_id = ?", id, ctx.Value("tenant_id").(int)).Delete(&domain.ProductVariant{}).Error
}
