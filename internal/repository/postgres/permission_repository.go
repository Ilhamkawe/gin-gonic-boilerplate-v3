package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"gorm.io/gorm"
)

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) domain.PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) Create(ctx context.Context, p *domain.Permission) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *permissionRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Permission, error) {
	var p domain.Permission
	err := r.db.WithContext(ctx).Where("uuid = ?", id).First(&p).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return &p, nil
}

func (r *permissionRepository) Fetch(ctx context.Context, limit int, offset int) ([]domain.Permission, int64, error) {
	var permissions []domain.Permission
	var count int64

	if err := r.db.WithContext(ctx).Model(&domain.Permission{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&permissions).Error
	if err != nil {
		return nil, 0, err
	}

	return permissions, count, nil
}

func (r *permissionRepository) Update(ctx context.Context, p *domain.Permission) error {
	result := r.db.WithContext(ctx).Model(&domain.Permission{}).Where("uuid = ?", p.UUID).Updates(p)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (r *permissionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.Permission{}, "uuid = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}
