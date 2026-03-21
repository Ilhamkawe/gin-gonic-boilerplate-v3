package database

import (
	"github.com/kawe/warehouse_backend/internal/domain"
	"github.com/kawe/warehouse_backend/pkg/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	logger.Info("Running auto migration...")

	// Drop old foreign key constraints that might block type changes for CreatedBy/UpdatedBy/DeletedBy
	// These are typically named as fk_<table_name>_<relation_name>
	constraintsToDrop := []struct {
		table      string
		constraint string
	}{
		{"stock_movements", "fk_stock_movements_user"},
		{"stock_movements", "fk_stock_movements_created_by"},
		{"transactions", "fk_transactions_user"},
		{"transaction_details", "fk_transaction_details_user"},
		{"products", "fk_products_user"},
		{"merchants", "fk_merchants_user"},
		{"warehouses", "fk_warehouses_user"},
		{"role_permissions", "fk_role_permissions_user"},
		{"roles", "fk_roles_user"},
	}

	for _, c := range constraintsToDrop {
		logger.Info("Dropping constraint: " + c.constraint + " on table: " + c.table)
		err := db.Exec("ALTER TABLE " + c.table + " DROP CONSTRAINT IF EXISTS " + c.constraint).Error
		if err != nil {
			logger.Error(err, "Failed to drop constraint "+c.constraint)
		}
	}

	// Tambahkan domain entity lain di sini jika ada fitur baru
	err := db.AutoMigrate(
		// Entities without foreign keys
		&domain.User{},
		&domain.Role{},
		&domain.Permission{},
		&domain.Category{},
		&domain.Tenant{},

		// Entities dependent on the base entities
		&domain.UserTenant{},
		&domain.RolePermission{},
		&domain.UserAccess{},
		&domain.Warehouse{},
		&domain.Product{},

		// Entities with deeper dependencies
		&domain.Merchant{},
		&domain.WarehouseProduct{},

		// Entities with most dependencies
		&domain.StockMovement{},
		&domain.Transaction{},
		&domain.TransactionDetail{},

		// subscription
		&domain.Plan{},
		&domain.AppPayment{},
	)

	if err != nil {
		logger.Error(err, "Migration failed")
		return err
	}

	logger.Info("Migration completed successfully")
	return nil
}
