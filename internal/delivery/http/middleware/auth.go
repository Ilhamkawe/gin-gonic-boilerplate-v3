package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"github.com/kawe/warehouse_backend/internal/dto"
	"github.com/kawe/warehouse_backend/pkg/jwt"
	"github.com/kawe/warehouse_backend/pkg/response"
)

func TenantAuthorization(tenantUsecase domain.TenantUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(dto.UserProfileResponse)
		tenantUUID, _ := uuid.Parse(c.GetHeader("X-Tenant-UUID"))
		if tenantUUID == uuid.Nil {
			response.Error(c, http.StatusUnauthorized, "Tenant ID required", nil)
			c.Abort()
			return
		}

		authorized := false
		for _, v := range user.Tenants {
			if v.Tenant.UUID == tenantUUID {
				authorized = true
				break
			}
		}

		if !authorized {
			response.Error(c, http.StatusUnauthorized, "User is not authorized to access this tenant", nil)
			c.Abort()
			return
		}

		tenant, err := tenantUsecase.GetByUUID(c, tenantUUID)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid authorization token", err.Error())
			c.Abort()
			return
		}

		c.Set("tenant_id", tenant.ID)
		c.Set("tenant_uuid", tenant.UUID)

		c.Next()
	}
}

func TenantTokenMatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantUUID, _ := uuid.Parse(c.GetHeader("X-Tenant-UUID"))

		// sekarang ada di state tenant apa
		active_tenant := c.MustGet("claims").(*jwt.AuthCustomClaims).TenantUUID
		if active_tenant == uuid.Nil || active_tenant != tenantUUID {
			response.Error(c, http.StatusUnauthorized, "Invalid authorization token", nil)
			c.Abort()
			return
		}

		c.Next()

	}
}

func AuthMiddleware(jwtService jwt.JWTService, userUsecase domain.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			response.Error(c, http.StatusUnauthorized, "Authorization token required", nil)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")

		token, err := jwtService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, "Invalid authorization token", nil)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*jwt.AuthCustomClaims); ok && token.Valid {
			c.Set("user_uuid", claims.UserUUID)
			c.Set("claims", claims)
		}

		userUUID := c.MustGet("user_uuid").(uuid.UUID)
		user, err := userUsecase.GetByUUID(c, userUUID)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid authorization token", err.Error())
			c.Abort()
			return
		}

		c.Set("user", dto.FromUserProfile(*user))

		c.Next()
	}
}
