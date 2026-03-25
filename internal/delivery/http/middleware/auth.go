package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"github.com/kawe/warehouse_backend/internal/dto"
	"github.com/kawe/warehouse_backend/pkg/jwt"
	"github.com/kawe/warehouse_backend/pkg/response"
	"gorm.io/gorm"
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

		tenant, err := tenantUsecase.GetAuthorizedTenant(c, tenantUUID, user.ID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusForbidden, "User is not authorized to access this tenant", err.Error())
			c.Abort()
			return
		} else if err != nil {
			response.Error(c, http.StatusInternalServerError, "Internal server error", err.Error())
			c.Abort()
			return
		}

		c.Set("tenant_id", tenant.ID)
		c.Set("tenant_uuid", tenant.UUID)

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
