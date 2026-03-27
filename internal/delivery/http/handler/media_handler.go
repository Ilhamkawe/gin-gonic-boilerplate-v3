package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kawe/warehouse_backend/internal/domain"
	"github.com/kawe/warehouse_backend/internal/dto"
	"github.com/kawe/warehouse_backend/pkg/response"
)

type MediaHandler struct {
	mediaUseCase domain.MediaUseCase
}

func NewMediaHandler(mediaUseCase domain.MediaUseCase) *MediaHandler {
	return &MediaHandler{
		mediaUseCase: mediaUseCase,
	}
}

func (h *MediaHandler) UploadPhoto(c *gin.Context) {
	file, header, err := c.Request.FormFile("photo")
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Photo is required", err.Error())
		return
	}
	defer file.Close()

	url, err := h.mediaUseCase.UploadPhoto(c, file, header.Size, header.Header.Get("Content-Type"))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to upload photo", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Photo uploaded successfully", dto.MediaResponse{
		URL:      url,
		FileName: header.Filename,
	})
}
