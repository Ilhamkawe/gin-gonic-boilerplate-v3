package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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
		response.Error(c, http.StatusInternalServerError, "Internal Servr Error", err)
		return
	}

	tenantResponse := dto.FromTenant(tenantDomain)

	response.Success(c, http.StatusCreated, "Tenant created successfully", tenantResponse)
}

// func (t *TenantHandler) GetByID(c *gin.Context) {
// 	id := c.Param("id")
// 	if err := t.tenantUseCase.GetByID(c, uuid.Must(uuid.Parse(id))); err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, tenant)
// }

// func (t *TenantHandler) Fetch(c *gin.Context) {
// 	limit := c.Query("limit")
// 	offset := c.Query("offset")
// 	if err := t.tenantUseCase.Fetch(c, limit, offset); err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, tenant)
// }

// func (t *TenantHandler) Update(c *gin.Context) {
// 	id := c.Param("id")
// 	if err := t.tenantUseCase.Update(c, uuid.Must(uuid.Parse(id))); err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, tenant)
// }

// func (t *TenantHandler) Delete(c *gin.Context) {
// 	id := c.Param("id")
// 	if err := t.tenantUseCase.Delete(c, uuid.Must(uuid.Parse(id))); err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, gin.H{"message": "Tenant deleted successfully"})
// }

// func (t *TenantHandler) IsAuthroized(c *gin.Context) {
// 	id := c.Param("id")
// 	if err := t.tenantUseCase.IsAuthroized(c, uuid.Must(uuid.Parse(id))); err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, gin.H{"message": "Tenant authorized successfully"})
// }
