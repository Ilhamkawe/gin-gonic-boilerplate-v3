package domain

import (
	"context"
	"time"
)

type AuditLog struct {
	ID            int       `json:"id" gorm:"primaryKey;autoIncrement"`
	TenantID      int       `json:"tenant_id" gorm:"not null"`
	UserID        int       `json:"user_id" gorm:"not null"`
	AuditableType string    `json:"auditable_type" gorm:"not null"`
	AuditableID   string    `json:"auditable_id" gorm:"not null"`
	Event         string    `json:"event" gorm:"not null"`
	URL           string    `json:"url" gorm:"not null"`
	IPAddress     string    `json:"ip_address" gorm:"not null"`
	UserAgent     string    `json:"user_agent" gorm:"not null"`
	Created_at    time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Deleted_at    time.Time `json:"deleted_at" gorm:"autoDeleteTime"`
}

type AuditLogRepository interface {
	Create(ctx context.Context, auditLog *AuditLog) error
}

type AuditLogUsecase interface {
	Create(ctx context.Context, auditLog *AuditLog) error
}
