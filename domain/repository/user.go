package repository

import (
	"base-gin-golang/domain/entity"
	"context"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByID(ctx context.Context, id int) (*entity.User, error)
}
