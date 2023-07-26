package product

import (
	"base-gin-golang/domain/entity"

	"base-gin-golang/domain/repository"
)

type GetProductByIDInput struct {
	ID int64
}

func GetByID(productRepository repository.ProductRepository, input *GetProductByIDInput) (*entity.Product, error) {
	product, err := productRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}
