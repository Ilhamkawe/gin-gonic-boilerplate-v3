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

type UserTenantHandler struct {
	userTenantUsecase domain.UserTenantUseCase
	validator         *validator.CustomValidator
}

func NewUserTenantHandler(userTenantUsecase domain.UserTenantUseCase, v *validator.CustomValidator) *UserTenantHandler {
	return &UserTenantHandler{
		userTenantUsecase: userTenantUsecase,
		validator:         v,
	}
}

func (h *UserTenantHandler) Create(c *gin.Context) {
	var req dto.CreateUserTenantDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	userUUID := c.MustGet("user_uuid").(uuid.UUID).String()

	ut := domain.UserTenant{
		UserID:    req.UserID,
		TenantID:  req.TenantID,
		RoleID:    req.RoleID,
		CreatedBy: userUUID,
	}

	if err := h.userTenantUsecase.Create(c.Request.Context(), &ut); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create user tenant mapping", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User tenant mapping created successfully", ut)
}

func (h *UserTenantHandler) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	uts, total, err := h.userTenantUsecase.Fetch(c.Request.Context(), limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch user tenant mappings", err.Error())
		return
	}

	response.Paginate(c, http.StatusOK, "User tenant mappings fetched successfully", response.PaginatedData{
		Items:      uts,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: int((total + int64(limit) - 1) / int64(limit)),
	})
}

func (h *UserTenantHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	if err := h.userTenantUsecase.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete mapping", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Mapping deleted successfully", nil)
}
