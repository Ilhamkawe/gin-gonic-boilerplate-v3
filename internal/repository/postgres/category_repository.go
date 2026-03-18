package postgres

import (
	"context"

	"github.com/google/uuid"
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

func (r *categoryRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	var category domain.Category
	err := r.db.Where("uuid = ?", id).First(&category).Error
	return &category, err
}

func (r *categoryRepository) Fetch(ctx context.Context, limit int, offset int) ([]domain.Category, int64, error) {
	var categories []domain.Category
	err := r.db.Limit(limit).Offset(offset).Find(&categories).Error
	return categories, int64(len(categories)), err
}

func (r *categoryRepository) Update(ctx context.Context, category *domain.Category) error {
	err := r.db.Save(category).Error
	return err
}

func (r *categoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.db.Where("uuid = ?", id).Delete(&domain.Category{}).Error
	return err
}
