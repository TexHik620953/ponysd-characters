package filestorage

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"path/filepath"

	"github.com/minio/minio-go/v7"
)

const bucketName = "files"

type Service struct {
	minioClient *minio.Client
}

func New(minioClient *minio.Client) *Service {
	return &Service{
		minioClient: minioClient,
	}
}

func (svc *Service) PrepareBuckets(ctx context.Context) error {
	exists, err := svc.minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to check if bucket exists: %w", err)
	}

	if !exists {
		err = svc.minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("failed to create bucket '%s': %w", bucketName, err)
		}
	}
	return nil
}

func calculateName(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func (svc *Service) StoreFile(ctx context.Context, img []byte) (string, error) {
	fileHash := calculateName(img)

	dir1 := fileHash[0:2]
	dir2 := fileHash[2:4]
	dir3 := fileHash[4:6]

	fineName := filepath.Join(dir1, dir2, dir3, fileHash)

	_, err := svc.minioClient.PutObject(
		ctx,
		bucketName,
		fineName,
		bytes.NewReader(img),
		int64(len(img)),
		minio.PutObjectOptions{
			ContentType: "image/png",
		},
	)

	if err != nil {
		return "", fmt.Errorf("failed to store image in MinIO: %w", err)
	}

	return fileHash, nil
}

func (svc *Service) GetFile(ctx context.Context, fileHash string) ([]byte, error) {
	dir1 := fileHash[0:2]
	dir2 := fileHash[2:4]
	dir3 := fileHash[4:6]

	fineName := filepath.Join(dir1, dir2, dir3, fileHash)

	// Получаем объект из MinIO
	obj, err := svc.minioClient.GetObject(
		ctx,
		bucketName,
		fineName,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get object from MinIO: %w", err)
	}
	defer obj.Close()

	imageData, err := io.ReadAll(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to read image data: %w", err)
	}

	if len(imageData) == 0 {
		return nil, fmt.Errorf("image not found: %s", fileHash)
	}

	return imageData, nil
}
