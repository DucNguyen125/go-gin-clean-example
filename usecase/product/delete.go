package product

import (
	"base-gin-golang/pkg/logger"

	"github.com/gin-gonic/gin"
)

type DeleteProductInput struct {
	ID int64
}

type DeleteProductOutPut struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func (pu *productUseCase) Delete(ctx *gin.Context, input *DeleteProductInput) (*DeleteProductOutPut, error) {
	rowsAffected, err := pu.productRepository.Delete(ctx, input.ID)
	if err != nil {
		logger.LogHandler(ctx, err)
		return &DeleteProductOutPut{
			RowsAffected: 0,
		}, err
	}
	output := &DeleteProductOutPut{
		RowsAffected: rowsAffected,
	}
	return output, nil
}
