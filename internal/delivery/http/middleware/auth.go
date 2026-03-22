package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"github.com/kawe/warehouse_backend/pkg/jwt"
)

func TenantAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*domain.User)
		if len(user.UserTenant) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user is not authorized to access this tenant"})
			c.Abort()
			return
		}

		userJson, _ := json.MarshalIndent(user, "", "  ")
		logger.Debug("user: \n" + string(userJson))

		c.Next()
	}
}

func AuthMiddleware(jwtService jwt.JWTService, userUsecase domain.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization token required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")

		token, err := jwtService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*jwt.AuthCustomClaims); ok && token.Valid {
			c.Set("user_uuid", claims.UserUUID)
		}

		userUUID := c.MustGet("user_uuid").(uuid.UUID)
		user, err := userUsecase.GetByUUID(c, userUUID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization token"})
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
