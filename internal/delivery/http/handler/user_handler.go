package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
	"github.com/kawe/warehouse_backend/internal/dto"
	"github.com/kawe/warehouse_backend/pkg/jwt"
	"github.com/kawe/warehouse_backend/pkg/response"
	"github.com/kawe/warehouse_backend/pkg/validator"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
	validator   *validator.CustomValidator
	jwtService  jwt.JWTService
}

func NewUserHandler(uu domain.UserUsecase, v *validator.CustomValidator) *UserHandler {
	return &UserHandler{
		userUsecase: uu,
		validator:   v,
	}
}

func (h *UserHandler) Login(c *gin.Context) {
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

	token, err := h.userUsecase.GenerateToken(c.Request.Context(), user.UUID, user.TenantID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate token", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User logged in successfully", dto.LoginResponse{
		User: dto.UserResponse{
			ID:        user.ID,
			UUID:      user.UUID,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Token: token,
	})
}

func (h *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.Error(c, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	user := &domain.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}

	if err := h.userUsecase.Create(c.Request.Context(), user); err != nil {
		if err == domain.ErrConflict {
			response.Error(c, http.StatusConflict, "User already exists", nil)
			return
		}
		response.Error(c, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User created successfully", dto.UserResponse{
		ID:        user.ID,
		UUID:      user.UUID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

func (h *UserHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	user, err := h.userUsecase.GetByID(c.Request.Context(), id)
	if err != nil {
		if err == domain.ErrNotFound {
			response.Error(c, http.StatusNotFound, "User not found", nil)
			return
		}
		response.Error(c, http.StatusInternalServerError, "Failed to fetch user", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User fetched successfully", dto.UserResponse{
		ID:        user.ID,
		UUID:      user.UUID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

func (h *UserHandler) Fetch(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	users, total, err := h.userUsecase.Fetch(c.Request.Context(), page, limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch users", err.Error())
		return
	}

	userResponses := make([]dto.UserResponse, 0)
	for _, user := range users {
		userResponses = append(userResponses, dto.UserResponse{
			ID:        user.ID,
			UUID:      user.UUID,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	response.Paginate(c, http.StatusOK, "Users fetched successfully", response.PaginatedData{
		Items:      userResponses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: int((total + int64(limit) - 1) / int64(limit)),
	})
}

func (h *UserHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	user := &domain.User{
		UUID:     id,
		Name:     req.Name,
		Password: req.Password,
	}

	if err := h.userUsecase.Update(c.Request.Context(), user); err != nil {
		if err == domain.ErrNotFound {
			response.Error(c, http.StatusNotFound, "User not found", nil)
			return
		}
		response.Error(c, http.StatusInternalServerError, "Failed to update user", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User updated successfully", nil)
}

func (h *UserHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	if err := h.userUsecase.Delete(c.Request.Context(), id); err != nil {
		if err == domain.ErrNotFound {
			response.Error(c, http.StatusNotFound, "User not found", nil)
			return
		}
		response.Error(c, http.StatusInternalServerError, "Failed to delete user", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User deleted successfully", nil)
}
