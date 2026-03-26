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

type RoleHandler struct {
	roleUsecase domain.RoleUsecase
	validator   *validator.CustomValidator
}

func NewRoleHandler(roleUsecase domain.RoleUsecase, v *validator.CustomValidator) *RoleHandler {
	return &RoleHandler{
		roleUsecase: roleUsecase,
		validator:   v,
	}
}

func (h *RoleHandler) Create(c *gin.Context) {
	var req dto.CreateRoleDTO
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

	role := domain.Role{
		UUID:      uuid.New(),
		Name:      req.Name,
		TenantID:  tenantID,
		CreatedBy: userUUID,
	}

	if err := h.roleUsecase.Create(c.Request.Context(), &role); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create role", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Role created successfully", dto.FromRole(role))
}

func (h *RoleHandler) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	roles, total, err := h.roleUsecase.Fetch(c.Request.Context(), limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch roles", err.Error())
		return
	}

	response.Paginate(c, http.StatusOK, "Roles fetched successfully", response.PaginatedData{
		Items:      dto.FromRoles(roles),
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: int((total + int64(limit) - 1) / int64(limit)),
	})
}

func (h *RoleHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid UUID", err.Error())
		return
	}

	role, err := h.roleUsecase.GetByID(c.Request.Context(), id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Role not found", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Role fetched successfully", dto.FromRole(*role))
}

func (h *RoleHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid UUID", err.Error())
		return
	}

	var req dto.UpdateRoleDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	userUUID := c.MustGet("user_uuid").(uuid.UUID).String()

	role := domain.Role{
		UUID:      id,
		Name:      req.Name,
		UpdatedBy: userUUID,
	}

	if err := h.roleUsecase.Update(c.Request.Context(), &role); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update role", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Role updated successfully", nil)
}

func (h *RoleHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid UUID", err.Error())
		return
	}

	if err := h.roleUsecase.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete role", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Role deleted successfully", nil)
}
