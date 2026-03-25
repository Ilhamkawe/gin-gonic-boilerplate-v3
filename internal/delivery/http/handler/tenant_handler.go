package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"github.com/kawe/warehouse_backend/internal/dto"
	"github.com/kawe/warehouse_backend/pkg/response"
	"github.com/kawe/warehouse_backend/pkg/validator"
)

type TenantHandler struct {
	tenantUseCase domain.TenantUseCase
	validator     *validator.CustomValidator
}

func NewTenantHandler(tenantUseCase domain.TenantUseCase, validator *validator.CustomValidator) *TenantHandler {
	return &TenantHandler{tenantUseCase: tenantUseCase, validator: validator}
}

func (t *TenantHandler) Create(c *gin.Context) {
	var tenant dto.CreateTenantDTO
	if err := c.ShouldBind(&tenant); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	user := c.MustGet("user").(dto.UserProfileResponse)

	if err := t.validator.Validate(tenant); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	exist, _ := t.tenantUseCase.IsSubdomainExist(c, tenant.Subdomain)
	if exist {
		response.Error(c, http.StatusBadRequest, "Subdomain already exists", nil)
		return
	}

	file, header, err := c.Request.FormFile("photo")
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	tenantDomain := domain.Tenant{
		Name:      tenant.Name,
		Address:   tenant.Address,
		Phone:     tenant.Phone,
		Email:     strings.ToLower(tenant.Email),
		Photo:     tenant.Photo,
		Subdomain: strings.ToLower(tenant.Subdomain),
		OwnerId:   user.ID,
		CreatedBy: user.UUID.String(),
	}

	if err := t.tenantUseCase.Create(c, &tenantDomain, file, header.Size); err != nil {
		response.Error(c, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	tenantResponse := dto.FromTenant(tenantDomain)

	response.Success(c, http.StatusCreated, "Tenant created successfully", tenantResponse)
}

func (t *TenantHandler) GetByID(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	tenant, err := t.tenantUseCase.GetByID(c, uuid.Must(uuid.Parse(id)))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get tenant", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Tenant fetched successfully", dto.FromTenant(*tenant))
}

func (t *TenantHandler) Fetch(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	tenants, total, err := t.tenantUseCase.Fetch(c, limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch tenants", err.Error())
		return
	}

	tenantResponses := make([]dto.TenantResponseDTO, 0)
	for _, tenant := range tenants {
		tenantResponses = append(tenantResponses, dto.FromTenant(tenant))
	}

	response.Paginate(c, http.StatusOK, "Tenants fetched successfully", response.PaginatedData{
		Items:      tenantResponses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: int((total + int64(limit) - 1) / int64(limit)),
	})
}

func (t *TenantHandler) Update(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	var req dto.UpdateTenantDTO
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	if err := t.validator.Validate(req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	user := c.MustGet("user").(dto.UserProfileResponse)

	file, header, _ := c.Request.FormFile("photo")
	var fileSize int64
	if header != nil {
		fileSize = header.Size
	}

	tenantDomain := domain.Tenant{
		UUID:      uuid.Must(uuid.Parse(id)),
		Name:      req.Name,
		Address:   req.Address,
		Phone:     req.Phone,
		Email:     strings.ToLower(req.Email),
		Subdomain: strings.ToLower(req.Subdomain),
		UpdatedBy: user.UUID.String(),
	}

	if err := t.tenantUseCase.Update(c, &tenantDomain, file, fileSize); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update tenant", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Tenant updated successfully", dto.FromTenant(tenantDomain))
}

func (t *TenantHandler) Delete(c *gin.Context) {
	id := c.Param("uuid")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "UUID is required", nil)
		return
	}

	if err := t.tenantUseCase.Delete(c, uuid.Must(uuid.Parse(id))); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete tenant", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Tenant deleted successfully", nil)
}
