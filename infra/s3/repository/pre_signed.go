package repository

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (sr *storageRepository) GenPreSignURL(key string) (string, error) {
	preSignClient := s3.NewPresignClient(sr.s3Client.Client)
	request, err := preSignClient.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(sr.cfg.AwsS3Bucket),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(sr.cfg.AwsS3PreSignDurationHour) * time.Hour
	})
	if err != nil {
		return "", err
	}
	return request.URL, nil
}

func (sr *storageRepository) GenPreSignURLWithOutTimestamp(key string) (string, error) {
	fileNameOld := strings.Split(key, "/")[len(strings.Split(key, "/"))-1]
	var fileName, extension string
	fileSplit := strings.Split(fileNameOld, ".")
	if len(fileSplit) > 1 {
		extension = fileSplit[len(fileSplit)-1]
		fileNameSplit := strings.Split(fileNameOld, "-")
		fileName = strings.Join(fileNameSplit[:len(fileNameSplit)-1], "-")
	}
	urlNew, err := url.Parse(fileName)
	if err != nil {
		return "", err
	}
	urlNew.RawQuery = urlNew.Query().Encode()
	// fileNameEncoded := url.QueryEscape(fileName)
	fileDownload := fmt.Sprintf(`attachment; filename ="%s.%s"`, urlNew, extension)
	preSignClient := s3.NewPresignClient(sr.s3Client.Client)
	request, err := preSignClient.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket:                     aws.String(sr.cfg.AwsS3Bucket),
		Key:                        aws.String(key),
		ResponseContentDisposition: &fileDownload,
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(sr.cfg.AwsS3PreSignDurationHour) * time.Hour
	})
	if err != nil {
		return "", err
	}
	return request.URL, nil
}
