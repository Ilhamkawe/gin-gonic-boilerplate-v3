package database

import (
	"github.com/kawe/warehouse_backend/internal/domain"
	"github.com/kawe/warehouse_backend/pkg/logger"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	logger.Info("Running auto migration...")

	// Tambahkan domain entity lain di sini jika ada fitur baru
	err := db.AutoMigrate(
		// Entities without foreign keys
		&domain.User{},
		&domain.Role{},
		&domain.Permission{},
		&domain.Category{},

		// Entities dependent on the base entities
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
	)

	if err != nil {
		logger.Error(err, "Migration failed")
		return err
	}

	logger.Info("Migration completed successfully")
	return nil
}
