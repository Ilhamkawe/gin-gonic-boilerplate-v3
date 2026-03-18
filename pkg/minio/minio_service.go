package minio

import (
	"context"
	"fmt"
	"io"

	"github.com/kawe/warehouse_backend/internal/domain"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type minioService struct {
	client     *minio.Client
	bucketName string
}

func NewMinioService(endpoint, accessKey, secretKey, bucketName string, useSSL bool) (domain.StorageService, error) {

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		return nil, err
	}

	return &minioService{
		client:     minioClient,
		bucketName: bucketName,
	}, nil

}

func (m *minioService) UploadFile(ctx context.Context, objectName string, reader io.Reader, objectSize int64, contentType string) (string, error) {
	_, err := m.client.PutObject(ctx, m.bucketName, objectName, reader, objectSize, minio.PutObjectOptions{
		ContentType: contentType,
	})

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://localhost:9000/%s/%s", m.bucketName, objectName), nil
}

func (m *minioService) DeleteFile(ctx context.Context, objectName string) error {
	err := m.client.RemoveObject(ctx, m.bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
