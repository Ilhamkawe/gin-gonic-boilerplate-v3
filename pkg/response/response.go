package response

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type PaginatedData struct {
	Items      interface{} `json:"items"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalPages int         `json:"total_pages"`
}

func Success(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, status int, message string, errors interface{}) {
	var errResponse interface{}

	if os.Getenv("APP_DEV_MODE") == "true" {
		errResponse = errors
	} else {
		if status == 422 {
			errResponse = errors
		} else {
			errResponse = nil
		}
	}

	if status >= 500 {
		log.Printf("[ERROR %d] %s: %v", status, message, errors)
	}

	c.JSON(status, Response{
		Success: false,
		Message: message,
		Errors:  errResponse,
	})
}

func Paginate(c *gin.Context, status int, message string, data PaginatedData) {
	c.JSON(status, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}
