package usecase

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"
	"github.com/kawe/warehouse_backend/internal/domain"
)

type mediaUseCase struct {
	storageService domain.StorageService
}

func NewMediaUseCase(storageService domain.StorageService) domain.MediaUseCase {
	return &mediaUseCase{
		storageService: storageService,
	}
}

func (u *mediaUseCase) UploadPhoto(ctx context.Context, reader io.Reader, size int64, contentType string) (string, error) {
	// Generate a unique filename
	ext := "jpg"
	if strings.Contains(contentType, "png") {
		ext = "png"
	}
	// Add other formats if needed

	fileName := fmt.Sprintf("temp/%s.%s", uuid.New().String(), ext)

	return u.storageService.UploadFile(ctx, fileName, reader, size, contentType)
}
