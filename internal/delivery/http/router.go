package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kawe/warehouse_backend/internal/delivery/http/handler"
	"github.com/kawe/warehouse_backend/internal/delivery/http/middleware"
	"github.com/kawe/warehouse_backend/internal/domain"
	"github.com/kawe/warehouse_backend/pkg/jwt"
	"github.com/kawe/warehouse_backend/pkg/response"
)

func NewRouter(userHandler *handler.UserHandler,
	categoryHandler *handler.CategoryHandler,
	jwtService jwt.JWTService,
	userUsecase domain.UserUsecase,
	tenantUsecase domain.TenantUseCase,
	tenantHandler *handler.TenantHandler,
	auditLogUsecase domain.AuditLogUsecase,
	authHandler *handler.AuthorizationHandler,
) *gin.Engine {
	router := gin.New()

	router.Use(middleware.Logger())
	router.Use(middleware.CORS())
	router.Use(gin.Recovery())

	router.GET("/health", func(c *gin.Context) {
		response.Success(c, http.StatusOK, "Service is healthy", nil)
	})

	api := router.Group("/api")
	v1 := api.Group("/v1")
	{

		auth := v1.Group("/auth")
		auth.Use(middleware.AutoMutationAudit(auditLogUsecase))
		{
			auth.POST("/login", authHandler.Login)

			protectedRoute := auth.Use(
				middleware.AuthMiddleware(jwtService, userUsecase),
				middleware.TenantAuthorization(tenantUsecase))
			{
				protectedRoute.POST("/tenant", authHandler.AuthorizationToTenant)
			}
		}

		users := v1.Group("/users")
		users.Use(middleware.AuthMiddleware(jwtService, userUsecase))
		{
			users.POST("", userHandler.Create)
			users.GET("", userHandler.Fetch)
			users.GET("/profile", userHandler.GetProfile)
			users.GET("/:id", userHandler.GetByID)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
			users.GET("/debug", userHandler.Debug)
		}

		categories := v1.Group("/categories")
		categories.Use(
			middleware.AuthMiddleware(jwtService, userUsecase),
			middleware.TenantAuthorization(tenantUsecase),
			middleware.TenantTokenMatch(),
			middleware.AutoMutationAudit(auditLogUsecase))
		{
			categories.POST("", categoryHandler.Create)
			categories.PUT("/:uuid", categoryHandler.Update)
			categories.DELETE("/:uuid", categoryHandler.Delete)
			categories.GET("/insight", categoryHandler.GetInsight)
			categories.GET("/product-counts", categoryHandler.GetWithProductCount)
			categories.GET("", categoryHandler.Index)
			categories.GET("/:uuid", categoryHandler.GetByID)
		}

		tanants := v1.Group("/tenants")
		tanants.Use(
			middleware.AuthMiddleware(jwtService, userUsecase),
			middleware.AutoMutationAudit(auditLogUsecase))
		{
			tanants.POST("", tenantHandler.Create)
		}

	}

	return router
}
