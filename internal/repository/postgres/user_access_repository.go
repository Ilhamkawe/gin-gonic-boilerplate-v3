package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"gorm.io/gorm"
)

type userAccessRepository struct {
	db *gorm.DB
}

func NewUserAccessRepository(db *gorm.DB) domain.UserAccessRepository {
	return &userAccessRepository{db: db}
}

func (r *userAccessRepository) Create(ctx context.Context, ua *domain.UserAccess) error {
	return r.db.WithContext(ctx).Create(ua).Error
}

func (r *userAccessRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.UserAccess, error) {
	var ua domain.UserAccess
	err := r.db.WithContext(ctx).Where("uuid = ?", id).First(&ua).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return &ua, nil
}

func (r *userAccessRepository) Fetch(ctx context.Context, limit int, offset int) ([]domain.UserAccess, int64, error) {
	var userAccesses []domain.UserAccess
	var count int64

	if err := r.db.WithContext(ctx).Model(&domain.UserAccess{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&userAccesses).Error
	if err != nil {
		return nil, 0, err
	}

	return userAccesses, count, nil
}

func (r *userAccessRepository) Update(ctx context.Context, ua *domain.UserAccess) error {
	return r.db.WithContext(ctx).Save(ua).Error
}

func (r *userAccessRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Where("uuid = ?", id).Delete(&domain.UserAccess{}).Error
}
