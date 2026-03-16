package database

import (
	"fmt"

	"github.com/kawe/warehouse_backend/internal/infrastructure/config"
	"github.com/kawe/warehouse_backend/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConn(cfg config.Config) (*gorm.DB, error) {
	dbURL := cfg.DatabaseURL
	if dbURL == "" {
		dbURL = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Connection Pool Settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	logger.Info("Successfully connected to PostgreSQL via GORM")
	return db, nil
}
