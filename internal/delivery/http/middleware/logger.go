package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kawe/warehouse_backend/pkg/logger"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method

		if raw != "" {
			path = path + "?" + raw
		}

		logger.Info("HTTP Request | %d | %13v | %s | %-7s %s",
			status,
			latency,
			c.ClientIP(),
			method,
			path,
		)
	}
}
