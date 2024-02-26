package product

import (
	"context"

	"base-gin-golang/pkg/logger"
)

type DeleteProductInput struct {
	ID int64 `path:"id"`
}

type DeleteProductOutputBody struct {
	RowsAffected int64 `json:"rowsAffected"`
}

type DeleteProductOutput struct {
	Body DeleteProductOutputBody
}

func (pu *productUseCase) Delete(
	ctx context.Context,
	input *DeleteProductInput,
) (*DeleteProductOutput, error) {
	rowsAffected, err := pu.productRepository.Delete(ctx, input.ID)
	if err != nil {
		logger.LogHandler(ctx, err)
		return &DeleteProductOutput{
			Body: DeleteProductOutputBody{
				RowsAffected: 0,
			},
		}, err
	}
	return &DeleteProductOutput{
		Body: DeleteProductOutputBody{
			RowsAffected: rowsAffected,
		},
	}, nil
}
