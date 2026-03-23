package middleware

import (
	"context"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kawe/warehouse_backend/internal/domain"
	"github.com/kawe/warehouse_backend/internal/dto"
)

func generateAuditType(c *gin.Context) string {

	path := c.FullPath()
	p := strings.TrimPrefix(path, "/api/v1/")

	parts := strings.Split(p, "/")

	var cleanParts []string

	for _, part := range parts {
		if !strings.HasPrefix(part, ":") &&
			!strings.HasPrefix(part, "*") &&
			part != "" {
			cleanParts = append(cleanParts, strings.ToUpper(part))
		}
	}

	return cleanParts[0]
}

func generateActionName(method, path string, c *gin.Context) string {
	fullPath := c.FullPath()

	exceptions := map[string]string{
		"POST/api/v1/auth/login": "LOGIN",
	}

	if exception, ok := exceptions[method+fullPath]; ok {
		return exception
	}

	p := strings.TrimPrefix(fullPath, "/api/v1/")

	parts := strings.Split(p, "/")

	var cleanParts []string

	for _, part := range parts {
		if !strings.HasPrefix(part, ":") &&
			!strings.HasPrefix(part, "*") &&
			part != "" {
			cleanParts = append(cleanParts, strings.ToUpper(part))
		}
	}

	resource := strings.Join(cleanParts, "_")

	mapping := map[string]string{
		"POST":   "CREATE_",
		"PUT":    "UPDATE_",
		"PATCH":  "MODIFY_",
		"DELETE": "DELETE_",
	}

	return mapping[method] + resource
}

func AutoMutationAudit(auditLogUsecase domain.AuditLogUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		isMutation := c.Request.Method == "POST" ||
			c.Request.Method == "PUT" ||
			c.Request.Method == "PATCH" ||
			c.Request.Method == "DELETE"

		c.Next()

		if !isMutation {

			return
		}

		actionName := generateActionName(c.Request.Method, c.FullPath(), c)

		path := c.Request.URL.Path
		ip := c.ClientIP()
		ua := c.Request.UserAgent()
		tenantIDHeader := c.GetHeader("X-Tenant-ID")

		var userID int
		var tenantID int
		var auditableID string
		var auditableType string

		if t, exists := c.Get("tenant_id"); exists {
			tenantID = t.(int)
		} else if tenantIDHeader != "" {
			tenantID, _ = strconv.Atoi(tenantIDHeader)
		}

		if u, exists := c.Get("user"); exists {
			if user, ok := u.(dto.UserProfileResponse); ok {
				userID = user.ID
				auditableID = user.UUID.String()
				auditableType = "User"

				if tenantID == 0 && len(user.Tenants) > 0 {
					tenantID = user.Tenants[0].ID
				}
			}
		}

		auditableType = generateAuditType(c)

		go func() {
			auditLogUsecase.Create(context.Background(), &domain.AuditLog{
				TenantID:      tenantID,
				UserID:        userID,
				AuditableType: auditableType,
				AuditableID:   auditableID,
				Event:         actionName,
				URL:           path,
				IPAddress:     ip,
				UserAgent:     ua,
			})
		}()

	}
}
