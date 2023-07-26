package repository

import (
	"base-gin-golang/domain/entity"
	"base-gin-golang/domain/repository"
	"base-gin-golang/infra/postgresql"
	"base-gin-golang/infra/postgresql/model"
	"context"

	dataPkg "base-gin-golang/pkg/data"

	"gorm.io/gorm"
)

type productRepository struct {
	db          *postgresql.Database
	dataService dataPkg.Service
}

func NewProductRepository(db *postgresql.Database, dataService dataPkg.Service) repository.ProductRepository {
	return &productRepository{
		db:          db,
		dataService: dataService,
	}
}

func (r *productRepository) Create(ctx context.Context, input *entity.Product) (*entity.Product, error) {
	query := r.db.WithContext(ctx)
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if ok {
		query = tx.WithContext(ctx)
	}
	product := &model.Product{}
	err := r.dataService.Copy(product, input)
	if err != nil {
		return nil, err
	}
	result := query.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	err = r.dataService.Copy(input, product)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (r *productRepository) GetList(ctx context.Context, input entity.GetListProductOption) ([]*entity.Product, error) {
	result := []*entity.Product{}
	query := r.db.WithContext(ctx).Model(&model.Product{})
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

func (r *productRepository) GetByID(ctx context.Context, id int64) (*entity.Product, error) {
	result := &entity.Product{}
	err := r.db.WithContext(ctx).Model(&model.Product{}).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *productRepository) Update(ctx context.Context, id int64, input *entity.Product) (*entity.Product, error) {
	query := r.db.WithContext(ctx)
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if ok {
		query = tx.WithContext(ctx)
	}
	product := &model.Product{}
	err := r.dataService.Copy(product, input)
	if err != nil {
		return nil, err
	}
	result := query.Model(&model.Product{}).Where("id = ?", id).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	err = r.dataService.Copy(input, product)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (r *productRepository) Delete(ctx context.Context, id int64) (int64, error) {
	query := r.db.WithContext(ctx)
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if ok {
		query = tx.WithContext(ctx)
	}
	result := query.Delete(&model.Product{}, id)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
