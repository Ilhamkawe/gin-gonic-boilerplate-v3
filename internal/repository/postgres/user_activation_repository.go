package postgres

import (
	"context"

	"github.com/kawe/warehouse_backend/internal/domain"
	"gorm.io/gorm"
)

type userActivationRepository struct {
	db *gorm.DB
}

func NewUserActivationRepository(db *gorm.DB) domain.UserActivationRepository {
	return &userActivationRepository{db: db}
}

func (r *userActivationRepository) Create(ctx context.Context, activation *domain.UserActivation) error {
	return r.db.Create(activation).Error
}

func (r *userActivationRepository) GetByToken(ctx context.Context, token string) (*domain.UserActivation, error) {
	var activation domain.UserActivation
	if err := r.db.Where("token = ? AND is_used = ?", token, false).First(&activation).Error; err != nil {
		return nil, err
	}
	return &activation, nil
}

func (r *userActivationRepository) Update(ctx context.Context, activation *domain.UserActivation) error {
	return r.db.Save(activation).Error
}
