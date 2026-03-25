package postgres

import (
	"context"

	"github.com/kawe/warehouse_backend/internal/domain"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(ctx context.Context, category *domain.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetByID(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	err := r.db.Where("uuid = ? AND tenant_id = ?", category.UUID, category.TenantID).First(category).Error
	return category, err
}

func (r *categoryRepository) Fetch(ctx context.Context, limit int, offset int) ([]domain.Category, int64, error) {
	var categories []domain.Category
	var total int64
	tenantID := ctx.Value("tenant_id").(int)

	if err := r.db.Model(&domain.Category{}).Where("tenant_id = ? AND deleted_at IS NULL", tenantID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Where("tenant_id = ? AND deleted_at IS NULL", tenantID).Limit(limit).Offset(offset).Find(&categories).Error
	return categories, total, err
}

func (r *categoryRepository) Update(ctx context.Context, category *domain.Category) error {
	err := r.db.Where("uuid = ? AND tenant_id = ?", category.UUID, category.TenantID).Updates(category).Error
	return err
}

func (r *categoryRepository) Delete(ctx context.Context, category *domain.Category) error {
	err := r.db.Where("uuid = ? AND tenant_id = ?", category.UUID, category.TenantID).Delete(&domain.Category{}).Error
	return err
}

func (r *categoryRepository) GetInsight(ctx context.Context) (*domain.InsightCategory, error) {
	var insight domain.InsightCategory
	err := r.db.Model(&domain.Category{}).Select("count(*) as total_categories, count(CASE WHEN deleted_at IS NULL THEN 1 END) as active_categories, count(CASE WHEN deleted_at IS NOT NULL THEN 1 END) as inactive_categories").Scan(&insight).Error
	return &insight, err
}

func (r *categoryRepository) FetchWithProductCount(ctx context.Context, tenantID int) ([]domain.CategoryWithCount, error) {
	var results []domain.CategoryWithCount
	err := r.db.Table("categories").
		Select("categories.*, COUNT(products.category_id) as product_count").
		Joins("LEFT JOIN products ON products.category_id = categories.id").
		Where("categories.tenant_id = ? AND categories.deleted_at IS NULL", tenantID).
		Group("categories.id").
		Scan(&results).Error
	return results, err
}
