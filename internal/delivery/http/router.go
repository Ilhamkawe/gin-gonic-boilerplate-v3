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
	warehouseHandler *handler.WarehouseHandler,
	merchantHandler *handler.MerchantHandler,
	productHandler *handler.ProductHandler,
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
			auth.POST("/register", authHandler.Register)

			protectedRoute := auth.Use(
				middleware.AuthMiddleware(jwtService, userUsecase),
				middleware.TenantAuthorization(tenantUsecase))
			{
				protectedRoute.POST("/tenant", authHandler.AuthorizationToTenant)
			}
		}

		users := v1.Group("/users")
		{
			users.POST("", userHandler.Create)
			users.Use(middleware.AuthMiddleware(jwtService, userUsecase))
			{
				users.GET("", userHandler.Fetch)
				users.GET("/profile", userHandler.GetProfile)
				users.GET("/:id", userHandler.GetByID)
				users.PUT("/:id", userHandler.Update)
				users.DELETE("/:id", userHandler.Delete)
				users.GET("/debug", userHandler.Debug)
			}
		}

		categories := v1.Group("/categories")
		categories.Use(
			middleware.AuthMiddleware(jwtService, userUsecase),
			middleware.TenantAuthorization(tenantUsecase),
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
			middleware.AuthMiddleware(jwtService, userUsecase))
		{

			auditGroup := tanants.Group("")
			auditGroup.Use(middleware.AutoMutationAudit(auditLogUsecase))
			{
				tanants.POST("", tenantHandler.Create)
				tanants.GET("", tenantHandler.Fetch)
				tanants.GET("/:uuid", tenantHandler.GetByID)

				tanants.DELETE("/:uuid", tenantHandler.Delete)
			}

			authTenantGoup := tanants.Group("")
			authTenantGoup.Use(middleware.TenantAuthorization(tenantUsecase))
			{
				authTenantGoup.PUT("/:uuid", tenantHandler.Update)
			}

		}

		warehouses := v1.Group("/warehouses")
		warehouses.Use(
			middleware.AuthMiddleware(jwtService, userUsecase),
			middleware.TenantAuthorization(tenantUsecase),
			middleware.AutoMutationAudit(auditLogUsecase))
		{
			warehouses.POST("", warehouseHandler.Create)
			warehouses.GET("", warehouseHandler.Index)
			warehouses.GET("/:uuid", warehouseHandler.GetByID)
			warehouses.PUT("/:uuid", warehouseHandler.Update)
			warehouses.DELETE("/:uuid", warehouseHandler.Delete)
		}

		merchants := v1.Group("/merchants")
		merchants.Use(
			middleware.AuthMiddleware(jwtService, userUsecase),
			middleware.TenantAuthorization(tenantUsecase),
			middleware.AutoMutationAudit(auditLogUsecase))
		{
			merchants.POST("", merchantHandler.Create)
			merchants.GET("", merchantHandler.Index)
			merchants.GET("/:uuid", merchantHandler.GetByID)
			merchants.PUT("/:uuid", merchantHandler.Update)
			merchants.DELETE("/:uuid", merchantHandler.Delete)
		}

		products := v1.Group("/products")
		products.Use(
			middleware.AuthMiddleware(jwtService, userUsecase),
			middleware.TenantAuthorization(tenantUsecase),
			middleware.AutoMutationAudit(auditLogUsecase))
		{
			products.POST("", productHandler.Create)
			products.GET("", productHandler.Index)
			products.GET("/:uuid", productHandler.GetByID)
			products.PUT("/:uuid", productHandler.Update)
			products.DELETE("/:uuid", productHandler.Delete)
		}

	}

	return router
}

// notee :
// middleware.AuthMiddleware(jwtService, userUsecase), untuk endpoint yang perelu token login
// middleware.TenantAuthorization(tenantUsecase), // untuk endpoint yang perlu cek apakah user punya akses ke tenant
// middleware.TenantTokenMatch(), // untuk endpoint yang perlu cek apakah tenant_id di token cocok dengan tenant_id di request
// middleware.AutoMutationAudit(auditLogUsecase) // untuk endpoint yang perlu insert audit log
