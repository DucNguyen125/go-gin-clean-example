package product

import (
	"context"

	"base-gin-golang/domain/entity"
	"base-gin-golang/pkg/logger"
)

type GetProductByIDInput struct {
	ID int64 `path:"id"`
}

type GetProductByIDOutput struct {
	Body *entity.Product
}

func (u *productUseCase) GetByID(
	ctx context.Context,
	input *GetProductByIDInput,
) (*GetProductByIDOutput, error) {
	product, err := u.productRepository.GetByID(ctx, input.ID)
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	return &GetProductByIDOutput{Body: product}, nil
}
