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
)

type ProductVariantHandler struct {
	variantUseCase domain.ProductVariantUseCase
	validator      *validator.CustomValidator
}

func NewProductVariantHandler(variantUseCase domain.ProductVariantUseCase, validator *validator.CustomValidator) *ProductVariantHandler {
	return &ProductVariantHandler{variantUseCase: variantUseCase, validator: validator}
}

func (h *ProductVariantHandler) Create(c *gin.Context) {
	var req dto.CreateProductVariantDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	userUUID := c.MustGet("user_uuid").(uuid.UUID).String()
	tenantID := c.MustGet("tenant_id").(int)

	variantDomain := domain.ProductVariant{
		ProductID:     req.ProductID,
		SKU:           req.SKU,
		AttributeJson: req.AttributeJson,
		Price:         req.Price,
		Barcode:       req.Barcode,
		TenantID:      tenantID,
		CreatedBy:     userUUID,
	}

	ctx := c.Request.Context()
	if err := h.variantUseCase.Create(ctx, &variantDomain); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create product variant", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Product variant created successfully", dto.FromProductVariant(variantDomain))
}

func (h *ProductVariantHandler) GetByID(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	variant, err := h.variantUseCase.GetByID(c.Request.Context(), uuid.Must(uuid.Parse(id)))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get product variant", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Product variant fetched successfully", dto.FromProductVariant(*variant))
}

func (h *ProductVariantHandler) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	variants, total, err := h.variantUseCase.Fetch(c.Request.Context(), limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch product variants", err.Error())
		return
	}

	variantResponses := dto.FromProductVariants(variants)

	response.Paginate(c, http.StatusOK, "Product variants fetched successfully", response.PaginatedData{
		Items:      variantResponses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: int((total + int64(limit) - 1) / int64(limit)),
	})
}

func (h *ProductVariantHandler) Update(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	var req dto.UpdateProductVariantDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	userUUID := c.MustGet("user_uuid").(uuid.UUID).String()
	tenantID := c.MustGet("tenant_id").(int)

	variantDomain := domain.ProductVariant{
		UUID:          uuid.Must(uuid.Parse(id)),
		SKU:           req.SKU,
		AttributeJson: req.AttributeJson,
		Price:         req.Price,
		Barcode:       req.Barcode,
		TenantID:      tenantID,
		UpdatedBy:     userUUID,
	}

	if err := h.variantUseCase.Update(c.Request.Context(), &variantDomain); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update product variant", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Product variant updated successfully", dto.FromProductVariant(variantDomain))
}

func (h *ProductVariantHandler) Delete(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	if err := h.variantUseCase.Delete(c.Request.Context(), uuid.Must(uuid.Parse(id))); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete product variant", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Product variant deleted successfully", nil)
}
