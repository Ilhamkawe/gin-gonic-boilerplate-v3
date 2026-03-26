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

type PermissionHandler struct {
	permissionUsecase domain.PermissionUseCase
	validator         *validator.CustomValidator
}

func NewPermissionHandler(permissionUsecase domain.PermissionUseCase, v *validator.CustomValidator) *PermissionHandler {
	return &PermissionHandler{
		permissionUsecase: permissionUsecase,
		validator:         v,
	}
}

func (h *PermissionHandler) Create(c *gin.Context) {
	var req dto.CreatePermissionDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	userUUID := c.MustGet("user_uuid").(uuid.UUID).String()

	p := domain.Permission{
		Name:        req.Name,
		Module:      req.Module,
		Description: req.Description,
		IsAddon:     req.IsAddon,
		AddonID:     req.AddonID,
		CreatedBy:   userUUID,
	}

	if err := h.permissionUsecase.Create(c.Request.Context(), &p); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create permission", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Permission created successfully", dto.FromPermission(p))
}

func (h *PermissionHandler) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	permissions, total, err := h.permissionUsecase.Fetch(c.Request.Context(), limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch permissions", err.Error())
		return
	}

	response.Paginate(c, http.StatusOK, "Permissions fetched successfully", response.PaginatedData{
		Items:      dto.FromPermissions(permissions),
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: int((total + int64(limit) - 1) / int64(limit)),
	})
}

func (h *PermissionHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid UUID", err.Error())
		return
	}

	p, err := h.permissionUsecase.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Permission not found", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Permission fetched successfully", dto.FromPermission(*p))
}

func (h *PermissionHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid UUID", err.Error())
		return
	}

	var req dto.UpdatePermissionDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	userUUID := c.MustGet("user_uuid").(uuid.UUID).String()

	p := domain.Permission{
		UUID:        id,
		Name:        req.Name,
		Module:      req.Module,
		Description: req.Description,
		IsAddon:     req.IsAddon,
		AddonID:     req.AddonID,
		UpdatedBy:   userUUID,
	}

	if err := h.permissionUsecase.Update(c.Request.Context(), &p); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update permission", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Permission updated successfully", nil)
}

func (h *PermissionHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid UUID", err.Error())
		return
	}

	if err := h.permissionUsecase.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete permission", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Permission deleted successfully", nil)
}
