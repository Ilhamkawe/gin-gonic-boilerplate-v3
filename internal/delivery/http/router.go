package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kawe/warehouse_backend/internal/delivery/http/handler"
	"github.com/kawe/warehouse_backend/internal/delivery/http/middleware"
	"github.com/kawe/warehouse_backend/pkg/response"
)

func NewRouter(userHandler *handler.UserHandler) *gin.Engine {
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
		users := v1.Group("/users")
		{
			users.POST("", userHandler.Create)
			users.GET("", userHandler.Fetch)
			users.GET("/:id", userHandler.GetByID)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}
	}

	return router
}
