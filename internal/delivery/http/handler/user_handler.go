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

type UserHandler struct {
	userUsecase domain.UserUsecase
	validator   *validator.CustomValidator
}

func NewUserHandler(uu domain.UserUsecase, v *validator.CustomValidator) *UserHandler {
	return &UserHandler{
		userUsecase: uu,
		validator:   v,
	}
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
		Phone:    req.Phone,
	}

	if err := h.userUsecase.Create(c.Request.Context(), user); err != nil {
		if err == domain.ErrConflict {
			response.Error(c, http.StatusConflict, "User already exists", nil)
			return
		}
		response.Error(c, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User created successfully", dto.FromUser(*user))
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

	response.Success(c, http.StatusOK, "User fetched successfully", dto.FromUser(*user))
}

func (h *UserHandler) Fetch(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	users, total, err := h.userUsecase.Fetch(c.Request.Context(), offset, limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch users", err.Error())
		return
	}

	userResponses := dto.FromUsers(users)

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

func (h *UserHandler) Debug(c *gin.Context) {
	uuid := c.MustGet("user_uuid").(uuid.UUID)
	// logger.Debug("User UUID", uuid)
	user, err := h.userUsecase.GetByUUID(c.Request.Context(), uuid)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch user", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User fetched successfully", dto.FromUser(*user))
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userUUID, exists := c.Get("user_uuid")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	uuidParams := userUUID.(uuid.UUID)
	user, err := h.userUsecase.GetByUUID(c.Request.Context(), uuidParams)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch user profile", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User profile fetched successfully", dto.FromUserProfile(*user))
}
