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
		&domain.User{},
		&domain.Role{},
		&domain.Permission{},
		&domain.RolePermission{},
		&domain.UserAccess{},
	)
	
	if err != nil {
		logger.Error(err, "Migration failed")
		return err
	}
	
	logger.Info("Migration completed successfully")
	return nil
}
