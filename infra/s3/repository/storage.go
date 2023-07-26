package repository

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (sr *storageRepository) UploadFile(file *strings.Reader, filename string) (string, error) {
	objectInput := &s3.PutObjectInput{
		Bucket: aws.String(sr.cfg.AwsS3Bucket),
		Key:    aws.String(filename),
		Body:   file,
	}
	uploader := manager.NewUploader(sr.s3Client)
	result, err := uploader.Upload(context.Background(), objectInput)
	if err != nil {
		return "", err
	}
	return result.Location, nil
}
