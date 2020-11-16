package api

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/storage"
)

// GCPSvc GCP Service interface
type GCPSvc interface {
	CreateFile(ctx context.Context, name, content string) error
}

// GCPService GCP Service struct
type GCPService struct {
	storageClient *storage.Client
}

// NewGCPService inits gcp service
func NewGCPService() (*GCPService, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("init gcp storage client: %v", err)
	}

	svc := &GCPService{
		storageClient: client,
	}

	return svc, nil
}

const maxBytes = 1024

func (svc *GCPService) CreateFile(ctx context.Context, name, content string) error {
	bucketName := os.Getenv("BUCKET_NAME")
	obj := svc.storageClient.Bucket(bucketName).Object(name)
	w := obj.NewWriter(ctx)

	contentBytes := []byte(content)
	if len(contentBytes) > maxBytes {
		// trucate to max bytes
		contentBytes = contentBytes[:maxBytes]
	}

	_, err := w.Write(contentBytes)
	if err != nil {
		return fmt.Errorf("write: %v", err)
	}

	if err := w.Close(); err != nil {
		return fmt.Errorf("close: %v", err)
	}

	return nil
}
