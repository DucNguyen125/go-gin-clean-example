package product

import (
	"base-gin-golang/domain/entity"

	"base-gin-golang/domain/repository"
)

type GetProductByIdInput struct {
	Id int64
}

func GetById(productRepository repository.ProductRepository, input *GetProductByIdInput) (*entity.Product, error) {
	product, err := productRepository.GetById(input.Id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
