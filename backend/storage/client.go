package storage

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"
)

type S3Client struct {
	Client     *s3.Client
	BucketName string
}

// a function to upload a file to s3
func (s *S3Client) UploadFile(ctx context.Context, fileName string, data io.Reader) error {
	_, putErr := s.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &s.BucketName,
		Key:    aws.String(fileName),
		Body:   data,
	})
	if putErr != nil {
		log.Error().Err(putErr).Msg("Error uploading file")
		return putErr
	}
	return nil

}

// a function to get all files from a folder
func (s *S3Client) GetFiles(ctx context.Context, folderName string) ([]string, error) {
	var files []string
	resp, listErr := s.Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: &s.BucketName,
		Prefix: aws.String(folderName),
	})
	if listErr != nil {
		log.Error().Err(listErr).Msg("Error listing objects")
		return nil, listErr
	}
	for _, item := range resp.Contents {
		files = append(files, *item.Key)
	}
	return files, nil
}

// a function to Download a file from s3
func (s *S3Client) DownloadFile(ctx context.Context, filePath string) (io.ReadCloser, error) {
	resp, getErr := s.Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &s.BucketName,
		Key:    aws.String(filePath),
	})
	if getErr != nil {
		log.Error().Err(getErr).Msg("Error getting file")
		return nil, getErr
	}
	return resp.Body, nil
}
