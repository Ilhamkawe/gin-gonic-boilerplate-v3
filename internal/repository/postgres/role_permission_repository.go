package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"gorm.io/gorm"
)

type rolePermissionRepository struct {
	db *gorm.DB
}

func NewRolePermissionRepository(db *gorm.DB) domain.RolePermissionRepository {
	return &rolePermissionRepository{db: db}
}

func (r *rolePermissionRepository) Create(ctx context.Context, rp *domain.RolePermission) error {
	return r.db.WithContext(ctx).Create(rp).Error
}

func (r *rolePermissionRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.RolePermission, error) {
	var rp domain.RolePermission
	err := r.db.WithContext(ctx).Preload("Permission").Where("uuid = ?", id).First(&rp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return &rp, nil
}

func (r *rolePermissionRepository) Fetch(ctx context.Context, limit int, offset int) ([]domain.RolePermission, int64, error) {
	var rps []domain.RolePermission
	var count int64

	if err := r.db.WithContext(ctx).Model(&domain.RolePermission{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.WithContext(ctx).Preload("Permission").Limit(limit).Offset(offset).Find(&rps).Error
	if err != nil {
		return nil, 0, err
	}

	return rps, count, nil
}

func (r *rolePermissionRepository) Update(ctx context.Context, rp *domain.RolePermission) error {
	return r.db.WithContext(ctx).Save(rp).Error
}

func (r *rolePermissionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Where("uuid = ?", id).Delete(&domain.RolePermission{}).Error
}

func (r *rolePermissionRepository) BulkInsert(ctx context.Context, rps []domain.RolePermission) error {
	return r.db.WithContext(ctx).Create(&rps).Error
}
