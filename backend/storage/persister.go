package storage

import (
	"context"
	"io"
)

type S3Interface interface {
	UploadFile(ctx context.Context, fileName string, data io.Reader) error
	GetFiles(ctx context.Context, folderName string) ([]string, error)
	DownloadFile(ctx context.Context, filePath string) (io.ReadCloser, error)
}
