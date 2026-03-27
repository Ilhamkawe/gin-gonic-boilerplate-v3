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

type ProductHandler struct {
	productUseCase domain.ProductUseCase
	validator      *validator.CustomValidator
}

func NewProductHandler(productUseCase domain.ProductUseCase, validator *validator.CustomValidator) *ProductHandler {
	return &ProductHandler{productUseCase: productUseCase, validator: validator}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req dto.CreateProductDTO
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

	productDomain := domain.Product{
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		IsPopular:   req.IsPopular,
		Thumbnail:   req.Thumbnail,
		TenantID:    tenantID,
		CreatedBy:   userUUID,
	}

	if len(req.Variants) > 0 {
		var variants []domain.ProductVariant
		for _, v := range req.Variants {
			variants = append(variants, domain.ProductVariant{
				SKU:           v.SKU,
				AttributeJson: v.AttributeJson,
				Price:         v.Price,
				Barcode:       v.Barcode,
				TenantID:      tenantID,
				CreatedBy:     userUUID,
			})
		}
		productDomain.Variants = variants
	}

	tenantUUID := c.MustGet("tenant_uuid").(uuid.UUID).String()

	if err := h.productUseCase.Create(c, &productDomain, tenantUUID); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create product", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Product created successfully", dto.FromProduct(productDomain))
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	product, err := h.productUseCase.GetByID(c, uuid.Must(uuid.Parse(id)))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get product", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Product fetched successfully", dto.FromProduct(*product))
}

func (h *ProductHandler) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	products, total, err := h.productUseCase.Fetch(c, limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch products", err.Error())
		return
	}

	productResponses := dto.FromProducts(products)

	response.Paginate(c, http.StatusOK, "Products fetched successfully", response.PaginatedData{
		Items:      productResponses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: int((total + int64(limit) - 1) / int64(limit)),
	})
}

func (h *ProductHandler) Update(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	var req dto.UpdateProductDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	userUUID := c.MustGet("user_uuid").(uuid.UUID).String()
	tenantID := c.MustGet("tenant_id").(int)

	productDomain := domain.Product{
		UUID:        uuid.Must(uuid.Parse(id)),
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		IsPopular:   req.IsPopular,
		Thumbnail:   req.Thumbnail,
		TenantID:    tenantID,
		UpdatedBy:   userUUID,
	}

	tenantUUID := c.MustGet("tenant_uuid").(uuid.UUID).String()

	if err := h.productUseCase.Update(c, &productDomain, tenantUUID); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update product", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Product updated successfully", dto.FromProduct(productDomain))
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	if err := h.productUseCase.Delete(c, uuid.Must(uuid.Parse(id))); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete product", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Product deleted successfully", nil)
}
