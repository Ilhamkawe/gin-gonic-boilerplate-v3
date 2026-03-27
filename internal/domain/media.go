package domain

import (
	"context"
	"io"
)

type MediaUseCase interface {
	UploadPhoto(ctx context.Context, reader io.Reader, size int64, contentType string) (string, error)
}
