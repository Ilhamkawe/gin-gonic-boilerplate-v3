package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	var product domain.Product
	err := r.db.Where("uuid = ? AND tenant_id = ?", id, ctx.Value("tenant_id").(int)).First(&product).Error
	return &product, err
}

func (r *productRepository) Fetch(ctx context.Context, limit int, offset int) ([]domain.Product, int64, error) {
	var products []domain.Product
	var count int64
	tenantID := ctx.Value("tenant_id").(int)

	if err := r.db.Model(&domain.Product{}).Where("tenant_id = ? AND deleted_at IS NULL", tenantID).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Where("tenant_id = ? AND deleted_at IS NULL", tenantID).Limit(limit).Offset(offset).Find(&products).Error
	return products, count, err
}

func (r *productRepository) Update(ctx context.Context, product *domain.Product) error {
	return r.db.Where("uuid = ? AND tenant_id = ?", product.UUID, product.TenantID).Updates(product).Error
}

func (r *productRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Where("uuid = ? AND tenant_id = ?", id, ctx.Value("tenant_id").(int)).Delete(&domain.Product{}).Error
}
