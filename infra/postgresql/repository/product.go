package repository

import (
	"base-gin-golang/domain/entity"
	"base-gin-golang/domain/repository"
	"base-gin-golang/infra/postgresql"
	"base-gin-golang/infra/postgresql/model"

	"github.com/jinzhu/copier"
)

type productRepository struct {
	db *postgresql.Database
}

func NewProductRepository(db *postgresql.Database) repository.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Create(input *entity.Product) (*entity.Product, error) {
	product := &model.Product{}
	err := copier.Copy(product, input)
	if err != nil {
		return nil, err
	}
	result := r.db.Create(product)
	if result.Error != nil {
		return nil, err
	}
	err = copier.Copy(input, product)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (r *productRepository) GetList(input entity.GetListProductOption) ([]*entity.Product, error) {
	result := []*entity.Product{}
	query := r.db.Model(&model.Product{})
	if input.PageIndex > 0 {
		offset := (input.PageIndex - 1) * input.PageSize
		query = query.Offset(offset).Limit(input.PageSize)
	} else {
		query = query.Limit(input.PageSize)
	}
	if input.Order != nil && *input.Order != "" {
		query = query.Order(*input.Order)
	}
	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *productRepository) GetByID(id int64) (*entity.Product, error) {
	result := &entity.Product{}
	err := r.db.Model(&model.Product{}).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *productRepository) Update(id int64, input *entity.Product) (*entity.Product, error) {
	product := &model.Product{}
	err := copier.Copy(product, input)
	if err != nil {
		return nil, err
	}
	result := r.db.Model(&model.Product{}).Where("id = ?", id).Updates(product)
	if result.Error != nil {
		return nil, err
	}
	err = copier.Copy(input, product)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (r *productRepository) Delete(id int64) (int64, error) {
	result := r.db.Delete(&model.Product{}, id)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
