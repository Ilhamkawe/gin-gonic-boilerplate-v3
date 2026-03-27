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

type WarehouseHandler struct {
	warehouseUseCase domain.WarehouseUseCase
	validator        *validator.CustomValidator
}

func NewWarehouseHandler(warehouseUseCase domain.WarehouseUseCase, validator *validator.CustomValidator) *WarehouseHandler {
	return &WarehouseHandler{warehouseUseCase: warehouseUseCase, validator: validator}
}

func (h *WarehouseHandler) Create(c *gin.Context) {
	var req dto.CreateWarehouseDTO
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

	warehouseDomain := domain.Warehouse{
		Name:      req.Name,
		Address:   req.Address,
		Phone:     req.Phone,
		Email:     req.Email,
		Photo:     req.Photo,
		TenantID:  tenantID,
		CreatedBy: userUUID,
	}

	if err := h.warehouseUseCase.Create(c, &warehouseDomain); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create warehouse", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Warehouse created successfully", dto.FromWarehouse(warehouseDomain))
}

func (h *WarehouseHandler) GetByID(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	warehouse, err := h.warehouseUseCase.GetByID(c, uuid.Must(uuid.Parse(id)))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get warehouse", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Warehouse fetched successfully", dto.FromWarehouse(*warehouse))
}

func (h *WarehouseHandler) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	warehouses, total, err := h.warehouseUseCase.Fetch(c, limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch warehouses", err.Error())
		return
	}

	warehouseResponses := dto.FromWarehouses(warehouses)

	response.Paginate(c, http.StatusOK, "Warehouses fetched successfully", response.PaginatedData{
		Items:      warehouseResponses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: int((total + int64(limit) - 1) / int64(limit)),
	})
}

func (h *WarehouseHandler) Update(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	var req dto.UpdateWarehouseDTO
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

	warehouseDomain := domain.Warehouse{
		UUID:      uuid.Must(uuid.Parse(id)),
		Name:      req.Name,
		Address:   req.Address,
		Phone:     req.Phone,
		Email:     req.Email,
		Photo:     req.Photo,
		TenantID:  tenantID,
		UpdatedBy: userUUID,
	}

	if err := h.warehouseUseCase.Update(c, &warehouseDomain); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update warehouse", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Warehouse updated successfully", dto.FromWarehouse(warehouseDomain))
}

func (h *WarehouseHandler) Delete(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	if err := h.warehouseUseCase.Delete(c, uuid.Must(uuid.Parse(id))); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete warehouse", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Warehouse deleted successfully", nil)
}
