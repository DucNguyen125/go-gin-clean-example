package s3

import (
	"base-gin-golang/config"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	*s3.Client
}

func NewS3Client(cfg *config.Environment) (*Client, error) {
	s3config, err := awsConfig.LoadDefaultConfig(
		context.Background(),
		awsConfig.WithRegion(cfg.AwsRegion),
		awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(cfg.AwsAccessKeyID, cfg.AwsSecretAccessKey, ""),
		))
	if err != nil {
		return nil, err
	}
	s3Client := s3.NewFromConfig(s3config)
	params := &s3.ListObjectsInput{
		Bucket:  aws.String(cfg.AwsS3Bucket),
		MaxKeys: 1,
	}
	_, err = s3Client.ListObjects(context.Background(), params)
	if err != nil {
		return nil, err
	}
	return &Client{
		Client: s3Client,
	}, nil
}
