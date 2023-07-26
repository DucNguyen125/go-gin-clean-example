package repository

import (
	"base-gin-golang/domain/entity"
)

type ProductRepository interface {
	Create(input *entity.Product) (*entity.Product, error)
	GetList(query entity.GetListProductOption) ([]*entity.Product, error)
	GetById(id int64) (*entity.Product, error)
	Update(id int64, input *entity.Product) (*entity.Product, error)
	Delete(id int64) (rowsAffected int64, err error)
}
