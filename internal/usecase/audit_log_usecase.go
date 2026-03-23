package usecase

import (
	"context"

	"github.com/kawe/warehouse_backend/internal/domain"
)

type auditLogUsecase struct {
	auditLogRepo domain.AuditLogRepository
}

func NewAuditLogUsecase(auditLogRepo domain.AuditLogRepository) domain.AuditLogUsecase {
	return &auditLogUsecase{auditLogRepo: auditLogRepo}
}

func (u *auditLogUsecase) Create(ctx context.Context, auditLog *domain.AuditLog) error {
	return u.auditLogRepo.Create(ctx, auditLog)
}
