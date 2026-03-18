package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kawe/warehouse_backend/internal/domain"
	"github.com/kawe/warehouse_backend/pkg/response"
	"github.com/kawe/warehouse_backend/pkg/validator"
	"gorm.io/datatypes"
)

type CategoryHandler struct {
	categoryUsecase domain.CategoryUseCase
	validator       *validator.CustomValidator
}

func NewCategoryHandler(categoryUsecase domain.CategoryUseCase, validator *validator.CustomValidator) *CategoryHandler {
	return &CategoryHandler{categoryUsecase: categoryUsecase, validator: validator}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var category domain.Category

	category.Name = c.PostForm("name")
	category.Tagline = c.PostForm("tagline")

	formJsonStr := c.PostForm("form_json")
	if formJsonStr != "" {
		category.FormJson = datatypes.JSON([]byte(formJsonStr))
	}

	// Lakukan validasi nama & tagline (karena c.PostForm tidak otomatis memicu validasi struct default)
	if category.Name == "" {
		response.Error(c, http.StatusBadRequest, "Name is required", nil)
		return
	}

	file, header, err := c.Request.FormFile("icon")
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := h.categoryUsecase.Create(c, &category, file, header.Size); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create category", err)
		return
	}

	response.Success(c, http.StatusCreated, "Category created successfully", category)
}
