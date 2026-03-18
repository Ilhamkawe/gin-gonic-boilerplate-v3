package domain

import (
	"context"
	"io"
)

type StorageService interface {
	UploadFile(ctx context.Context, objectName string, reader io.Reader, objectSize int64, contentType string) (string, error)
	DeleteFile(ctx context.Context, objectName string) error
}
