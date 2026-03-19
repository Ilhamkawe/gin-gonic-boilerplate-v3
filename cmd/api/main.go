package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	api "github.com/kawe/warehouse_backend/internal/delivery/http"
	"github.com/kawe/warehouse_backend/internal/delivery/http/handler"
	"github.com/kawe/warehouse_backend/internal/infrastructure/config"
	"github.com/kawe/warehouse_backend/internal/infrastructure/database"
	"github.com/kawe/warehouse_backend/internal/repository/postgres"
	"github.com/kawe/warehouse_backend/internal/usecase"
	"github.com/kawe/warehouse_backend/pkg/jwt"
	"github.com/kawe/warehouse_backend/pkg/logger"
	"github.com/kawe/warehouse_backend/pkg/minio"
	"github.com/kawe/warehouse_backend/pkg/validator"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig(".")
	if err != nil {
		logger.Fatal(err, "Failed to load config")
	}

	if cfg.AppDevMode {
		os.Setenv("APP_DEV_MODE", "true")
	}

	// Connect to database
	db, err := database.NewPostgresConn(cfg)
	if err != nil {
		logger.Fatal(err, "Failed to connect to database")
	}

	// Auto Migration
	if err := database.Migrate(db); err != nil {
		logger.Fatal(err, "Failed to run migration")
	}

	// Close database connection on exit
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Initialize dependencies
	v := validator.NewCustomValidator()
	timeout := time.Duration(10) * time.Second
	jwtService := jwt.NewJWTService(cfg.JWTSecret)
	storageService, err := minio.NewMinioService(
		cfg.MinioEndpoint,
		cfg.MinioAccessKey,
		cfg.MinioSecretKey,
		cfg.MinioBucketName,
		cfg.MinioUseSSL,
	)
	if err != nil {
		logger.Fatal(err, "Failed to initialize storage service")
	}

	// User module
	userRepo := postgres.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, timeout, jwtService)
	userHandler := handler.NewUserHandler(userUsecase, v)

	// Category module
	categoryRepo := postgres.NewCategoryRepository(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo, storageService)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase, v)

	// Router
	r := api.NewRouter(userHandler, categoryHandler, jwtService)

	// Server
	if cfg.AppPort == "" {
		cfg.AppPort = "8080"
	}
	srv := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: r,
	}

	// Graceful shutdown
	go func() {
		logger.Info("Starting server on port %s", cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(err, "Failed to start server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal(err, "Server forced to shutdown")
	}

	logger.Info("Server exited gracefully")
}
