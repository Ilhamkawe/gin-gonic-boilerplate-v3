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

type RolePermissionHandler struct {
	rpUsecase domain.RolePermissionUseCase
	validator *validator.CustomValidator
}

func NewRolePermissionHandler(rpUsecase domain.RolePermissionUseCase, v *validator.CustomValidator) *RolePermissionHandler {
	return &RolePermissionHandler{
		rpUsecase: rpUsecase,
		validator: v,
	}
}

func (h *RolePermissionHandler) BulkAssign(c *gin.Context) {
	var req dto.AssignPermissionsDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	userUUID := c.MustGet("user_uuid").(uuid.UUID).String()
	tenantID := c.MustGet("tenant_id").(int)

	rps := make([]domain.RolePermission, 0)
	for _, pID := range req.PermissionIDs {
		rps = append(rps, domain.RolePermission{
			RoleID:       req.RoleID,
			PermissionID: pID,
			TenantID:     tenantID,
			CreatedBy:    userUUID,
		})
	}

	if err := h.rpUsecase.BulkInsert(c.Request.Context(), rps); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to assign permissions", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Permissions assigned successfully", nil)
}

func (h *RolePermissionHandler) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	rps, total, err := h.rpUsecase.Fetch(c.Request.Context(), limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch role permissions", err.Error())
		return
	}

	response.Paginate(c, http.StatusOK, "Role permissions fetched successfully", response.PaginatedData{
		Items:      dto.FromRolePermissions(rps),
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: int((total + int64(limit) - 1) / int64(limit)),
	})
}

func (h *RolePermissionHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid UUID", err.Error())
		return
	}

	if err := h.rpUsecase.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete role permission", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Role permission deleted successfully", nil)
}
