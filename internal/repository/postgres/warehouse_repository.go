package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"gorm.io/gorm"
)

type warehouseRepository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) *warehouseRepository {
	return &warehouseRepository{db: db}
}

func (r *warehouseRepository) Create(ctx context.Context, warehouse *domain.Warehouse) error {
	return r.db.Create(warehouse).Error
}

func (r *warehouseRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Warehouse, error) {
	var warehouse domain.Warehouse
	err := r.db.Where("uuid = ? AND tenant_id = ?", id, ctx.Value("tenant_id").(int)).First(&warehouse).Error
	return &warehouse, err
}

func (r *warehouseRepository) Fetch(ctx context.Context, limit int, offset int) ([]domain.Warehouse, int64, error) {
	var warehouses []domain.Warehouse
	err := r.db.Where("tenant_id = ? AND deleted_at IS NULL", ctx.Value("tenant_id").(int)).Limit(limit).Offset(offset).Find(&warehouses).Error
	return warehouses, int64(len(warehouses)), err
}

func (r *warehouseRepository) Update(ctx context.Context, warehouse *domain.Warehouse) error {
	err := r.db.Where("uuid = ? AND tenant_id = ?", warehouse.UUID, warehouse.TenantID).Updates(warehouse).Error
	return err
}

func (r *warehouseRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.db.Where("uuid = ? AND tenant_id = ?", id, ctx.Value("tenant_id").(int)).Delete(&domain.Warehouse{}).Error
	return err
}
