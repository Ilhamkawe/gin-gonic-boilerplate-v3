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

type UserAccessHandler struct {
	uaUsecase domain.UserAccessUseCase
	validator *validator.CustomValidator
}

func NewUserAccessHandler(uaUsecase domain.UserAccessUseCase, v *validator.CustomValidator) *UserAccessHandler {
	return &UserAccessHandler{
		uaUsecase: uaUsecase,
		validator: v,
	}
}

func (h *UserAccessHandler) Create(c *gin.Context) {
	var req dto.CreateUserAccessDTO
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

	ua := domain.UserAccess{
		UserID:       req.UserID,
		RoleID:       req.RoleID,
		EntityID:     req.EntityID,
		EntityType:   req.EntityType,
		TenantID:     tenantID,
		UserTenantID: req.UserTenantID,
		CreatedBy:    userUUID,
	}

	if err := h.uaUsecase.Create(c.Request.Context(), &ua); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create user access", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User access created successfully", dto.FromUserAccess(ua))
}

func (h *UserAccessHandler) Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	ua, total, err := h.uaUsecase.Fetch(c.Request.Context(), limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch user access", err.Error())
		return
	}

	response.Paginate(c, http.StatusOK, "User access list fetched successfully", response.PaginatedData{
		Items:      dto.FromUserAccesses(ua),
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: int((total + int64(limit) - 1) / int64(limit)),
	})
}

func (h *UserAccessHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid UUID", err.Error())
		return
	}

	if err := h.uaUsecase.Delete(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete user access", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User access deleted successfully", nil)
}
