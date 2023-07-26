package repository

import (
	"base-gin-golang/domain/entity"
	"context"
)

type ProductRepository interface {
	Create(ctx context.Context, input *entity.Product) (*entity.Product, error)
	GetList(ctx context.Context, query entity.GetListProductOption) ([]*entity.Product, error)
	GetByID(ctx context.Context, id int64) (*entity.Product, error)
	Update(ctx context.Context, id int64, input *entity.Product) (*entity.Product, error)
	Delete(ctx context.Context, id int64) (int64, error)
}
