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
	err := r.db.Limit(limit).Offset(offset).Find(&categories).Error
	return categories, int64(len(categories)), err
}

func (r *categoryRepository) Update(ctx context.Context, category *domain.Category) error {
	err := r.db.Where("uuid = ?", category.UUID).Updates(category).Error
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
