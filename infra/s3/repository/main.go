package repository

import (
	"base-gin-golang/config"
	"base-gin-golang/domain/repository"
	"base-gin-golang/infra/s3"
)

func NewStorageRepositoryRepository(
	cfg *config.Environment,
	s3Client *s3.Client,
) repository.StorageRepository {
	return &storageRepository{
		cfg:      cfg,
		s3Client: s3Client,
	}
}

type storageRepository struct {
	cfg      *config.Environment
	s3Client *s3.Client
}
