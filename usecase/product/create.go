package product

import (
	"context"

	"base-gin-golang/domain/entity"
	"base-gin-golang/pkg/logger"
)

type CreateProductInputBody struct {
	ProductCode string `json:"productCode" validate:"required"`
	ProductName string `json:"productName" validate:"required"`
	Price       int    `json:"price"       validate:"required"`
}

type CreateProductInput struct {
	Body CreateProductInputBody
}

type CreateProductOutput struct {
	Body *entity.Product
}

func (pu *productUseCase) Create(
	ctx context.Context,
	input *CreateProductInput,
) (*CreateProductOutput, error) {
	data := &entity.Product{}
	err := pu.dataService.Copy(data, &input.Body)
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	newProduct, err := pu.productRepository.Create(ctx, data)
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	return &CreateProductOutput{Body: newProduct}, nil
}
