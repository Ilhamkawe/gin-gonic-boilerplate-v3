package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"gorm.io/gorm"
)

type RoleRepo struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) domain.RoleRepository {
	return &RoleRepo{db: db}
}

func (t *RoleRepo) Create(ctx context.Context, role *domain.Role) error {
	return t.db.Create(role).Error
}

func (t *RoleRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Role, error) {
	var role domain.Role
	err := t.db.Where("uuid = ?", id).First(&role).Error
	return &role, err
}

func (t *RoleRepo) Fetch(ctx context.Context, limit int, offset int) ([]domain.Role, int64, error) {
	var roles []domain.Role
	err := t.db.Limit(limit).Offset(offset).Find(&roles).Error
	return roles, int64(len(roles)), err
}

func (t *RoleRepo) Update(ctx context.Context, role *domain.Role) error {
	return t.db.Save(role).Error
}

func (t *RoleRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return t.db.Where("uuid = ?", id).Delete(&domain.Role{}).Error
}
