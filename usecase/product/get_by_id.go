package product

import (
	"base-gin-golang/domain/entity"
)

type GetProductByIDInput struct {
	ID int64
}

func (pu *productUseCase) GetByID(input *GetProductByIDInput) (*entity.Product, error) {
	product, err := pu.productRepository.GetByID(input.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}
