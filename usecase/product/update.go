package product

import (
	"context"

	"base-gin-golang/domain/entity"
	"base-gin-golang/pkg/logger"

	"github.com/jinzhu/copier"
)

type UpdateProductInputBody struct {
	ProductCode string `json:"productCode" validate:"required"`
	ProductName string `json:"productName" validate:"required"`
	Price       int    `json:"price"       validate:"required"`
}

type UpdateProductInput struct {
	ID   int64 `path:"id"`
	Body UpdateProductInputBody
}

type UpdateProductOutput struct {
	Body *entity.Product
}

func (pu *productUseCase) Update(
	ctx context.Context,
	input *UpdateProductInput,
) (*UpdateProductOutput, error) {
	data := &entity.Product{}
	err := copier.Copy(data, &input.Body)
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	newProduct, err := pu.productRepository.Update(ctx, input.ID, data)
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	return &UpdateProductOutput{Body: newProduct}, nil
}
