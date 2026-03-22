package postgres

import (
	"context"

	"github.com/kawe/warehouse_backend/internal/domain"
	"gorm.io/gorm"
)

type userTenantRepository struct {
	db *gorm.DB
}

func NewUserTenantRepository(db *gorm.DB) domain.UserTenantRepository {
	return &userTenantRepository{db: db}
}

func (r *userTenantRepository) Create(ctx context.Context, userTenant *domain.UserTenant) error {
	return r.db.Create(userTenant).Error
}

func (r *userTenantRepository) Update(ctx context.Context, userTenant *domain.UserTenant) error {
	return r.db.Save(userTenant).Error
}

func (r *userTenantRepository) Delete(ctx context.Context, id int) error {
	return r.db.Delete(&domain.UserTenant{}, id).Error
}

func (r *userTenantRepository) GetByID(ctx context.Context, id int) (*domain.UserTenant, error) {
	var userTenant domain.UserTenant
	err := r.db.Preload("User").Preload("Tenant").Preload("Role").First(&userTenant, id).Error
	return &userTenant, err
}

func (r *userTenantRepository) GetAll(ctx context.Context) ([]domain.UserTenant, error) {
	var userTenants []domain.UserTenant
	err := r.db.Preload("User").Preload("Tenant").Preload("Role").Find(&userTenants).Error
	return userTenants, err
}

func (r *userTenantRepository) Fetch(ctx context.Context, limit int, offset int) ([]domain.UserTenant, int64, error) {
	var userTenants []domain.UserTenant
	var count int64
	err := r.db.Preload("User").Preload("Tenant").Preload("Role").Limit(limit).Offset(offset).Find(&userTenants).Count(&count).Error
	return userTenants, count, err
}
