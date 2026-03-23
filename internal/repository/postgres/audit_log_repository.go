package postgres

import (
	"context"

	"github.com/kawe/warehouse_backend/internal/domain"
	"gorm.io/gorm"
)

type AuditLogRepository struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) *AuditLogRepository {
	return &AuditLogRepository{db: db}
}

func (r *AuditLogRepository) Create(ctx context.Context, auditLog *domain.AuditLog) error {
	return r.db.WithContext(ctx).Create(auditLog).Error
}
