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

type MerchantHandler struct {
	merchantUseCase domain.MerchantUseCase
	validator       *validator.CustomValidator
}

func NewMerchantHandler(merchantUseCase domain.MerchantUseCase, validator *validator.CustomValidator) *MerchantHandler {
	return &MerchantHandler{merchantUseCase: merchantUseCase, validator: validator}
}

func (h *MerchantHandler) Create(c *gin.Context) {
	var req dto.CreateMerchantDTO
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

	merchantDomain := domain.Merchant{
		Name:         req.Name,
		Address:      req.Address,
		Phone:        req.Phone,
		Email:        req.Email,
		Photo:        req.Photo,
		KeeperId:     req.KeeperId,
		HasWarehouse: req.HasWarehouse,
		WarehouseId:  req.WarehouseId,
		TenantID:     tenantID,
		CreatedBy:    userUUID,
	}

	if err := h.merchantUseCase.Create(c, &merchantDomain); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create merchant", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Merchant created successfully", dto.FromMerchant(merchantDomain))
}

func (h *MerchantHandler) GetByID(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	merchant, err := h.merchantUseCase.GetByID(c, uuid.Must(uuid.Parse(id)))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get merchant", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Merchant fetched successfully", dto.FromMerchant(*merchant))
}

func (h *MerchantHandler) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	merchants, total, err := h.merchantUseCase.Fetch(c, limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch merchants", err.Error())
		return
	}

	merchantResponses := dto.FromMerchants(merchants)

	response.Paginate(c, http.StatusOK, "Merchants fetched successfully", response.PaginatedData{
		Items:      merchantResponses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: int((total + int64(limit) - 1) / int64(limit)),
	})
}

func (h *MerchantHandler) Update(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	var req dto.UpdateMerchantDTO
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

	merchantDomain := domain.Merchant{
		UUID:         uuid.Must(uuid.Parse(id)),
		Name:         req.Name,
		Address:      req.Address,
		Phone:        req.Phone,
		Email:        req.Email,
		Photo:        req.Photo,
		KeeperId:     req.KeeperId,
		HasWarehouse: req.HasWarehouse,
		WarehouseId:  req.WarehouseId,
		TenantID:     tenantID,
		UpdatedBy:    userUUID,
	}

	if err := h.merchantUseCase.Update(c, &merchantDomain); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update merchant", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Merchant updated successfully", dto.FromMerchant(merchantDomain))
}

func (h *MerchantHandler) Delete(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	if err := h.merchantUseCase.Delete(c, uuid.Must(uuid.Parse(id))); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete merchant", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Merchant deleted successfully", nil)
}
