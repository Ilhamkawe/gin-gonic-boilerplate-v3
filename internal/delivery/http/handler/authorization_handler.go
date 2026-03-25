package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"github.com/kawe/warehouse_backend/internal/dto"
	"github.com/kawe/warehouse_backend/pkg/jwt"
	"github.com/kawe/warehouse_backend/pkg/response"
	"github.com/kawe/warehouse_backend/pkg/validator"
)

type AuthorizationHandler struct {
	jwtService  jwt.JWTService
	userUsecase domain.UserUsecase
	validator   *validator.CustomValidator
}

func NewAuthorizationHandler(js jwt.JWTService, uu domain.UserUsecase, v *validator.CustomValidator) *AuthorizationHandler {
	return &AuthorizationHandler{jwtService: js, userUsecase: uu, validator: v}
}

func (h *AuthorizationHandler) AuthorizationToTenant(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt.AuthCustomClaims)
	tenantUUID := c.MustGet("tenant_uuid").(uuid.UUID)

	token, err := h.userUsecase.GenerateToken(c, claims.UserUUID, tenantUUID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate token", err.Error())
		return
	}

	user, err := h.userUsecase.GetByUUID(c.Request.Context(), claims.UserUUID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch user", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Tenant Access Granted", dto.LoginResponse{
		User:  dto.FromUser(*user),
		Token: token,
	})

}

func (h *AuthorizationHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	user, err := h.userUsecase.Register(c.Request.Context(), req.Email, req.Password, req.Name, req.Phone)
	if err != nil {
		if err == domain.ErrConflict {
			response.Error(c, http.StatusConflict, "User with this email already exists", nil)
			return
		}
		response.Error(c, http.StatusInternalServerError, "Failed to register", err.Error())
		return
	}

	token, err := h.userUsecase.GenerateToken(c.Request.Context(), user.UUID, uuid.Nil)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate token", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User registered successfully", dto.LoginResponse{
		User:  dto.FromUser(*user),
		Token: token,
	})
}

func (h *AuthorizationHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	user, err := h.userUsecase.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		if err == domain.ErrNotFound {
			response.Error(c, http.StatusNotFound, "User not found", nil)
			return
		}
		if err == domain.ErrUnauthorized {
			response.Error(c, http.StatusUnauthorized, "Invalid credentials", nil)
			return
		}
		response.Error(c, http.StatusInternalServerError, "Failed to login", err.Error())
		return
	}

	token, err := h.userUsecase.GenerateToken(c.Request.Context(), user.UUID, uuid.Nil)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate token", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User logged in successfully", dto.LoginResponse{
		User:  dto.FromUser(*user),
		Token: token,
	})
}
