package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"github.com/kawe/warehouse_backend/internal/dto"
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

func (h *CategoryHandler) Delete(c *gin.Context) {

	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	if err := h.categoryUsecase.Delete(c, uuid.Must(uuid.Parse(id))); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete category", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Category deleted successfully", nil)
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var category dto.CreateCategory

	category.TenantID = c.MustGet("tenant_id").(int)
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

	user_uuid := c.MustGet("user_uuid").(uuid.UUID).String()

	categoryDomain := domain.Category{
		Name:      category.Name,
		Tagline:   category.Tagline,
		FormJson:  category.FormJson,
		TenantID:  category.TenantID,
		CreatedBy: user_uuid,
	}

	if err := h.categoryUsecase.Create(c, &categoryDomain, file, header.Size); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create category", err)
		return
	}

	response.Success(c, http.StatusCreated, "Category created successfully", category)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	var category dto.UpdateCategory

	// read value from queryparam
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	category.UUID = uuid.Must(uuid.Parse(id))

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

	user_uuid := c.MustGet("user_uuid").(uuid.UUID).String()

	categoryDomain := domain.Category{
		UUID:      category.UUID,
		Name:      category.Name,
		Tagline:   category.Tagline,
		FormJson:  category.FormJson,
		UpdatedBy: user_uuid,
	}

	if err := h.categoryUsecase.Update(c, &categoryDomain, file, header.Size); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update category", err)
		return
	}

	response.Success(c, http.StatusCreated, "Category updated successfully", category)
}

func (h *CategoryHandler) Index(c *gin.Context) {
	var categories []domain.Category

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	categories, total, err := h.categoryUsecase.Fetch(c, limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch category", err)
		return
	}

	categoryResponses := dto.FromCategories(categories)

	response.Paginate(c, http.StatusOK, "Category fetched successfully", response.PaginatedData{
		Items:      categoryResponses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: int((total + int64(limit) - 1) / int64(limit)),
	})
}

func (h *CategoryHandler) GetInsight(c *gin.Context) {
	insight, err := h.categoryUsecase.GetInsight(c)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get insight", err)
		return
	}

	response.Success(c, http.StatusOK, "Insight fetched successfully", dto.FromInsightCategory(*insight))
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	category, err := h.categoryUsecase.GetByID(c, uuid.Must(uuid.Parse(id)))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get category", err)
		return
	}

	response.Success(c, http.StatusOK, "Category fetched successfully", dto.FromCategory(*category))
}

func (h *CategoryHandler) GetWithProductCount(c *gin.Context) {
	tenantID := c.MustGet("tenant_id").(int)
	categories, err := h.categoryUsecase.FetchWithProductCount(c, tenantID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch categories with product count", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Categories with product count fetched successfully", dto.FromCategoriesWithProductCount(categories))
}
