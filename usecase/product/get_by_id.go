package product

import (
	"base-gin-golang/domain/entity"
	"base-gin-golang/pkg/logger"

	"github.com/gin-gonic/gin"
)

type GetProductByIDInput struct {
	ID int64
}

func (pu *productUseCase) GetByID(ctx *gin.Context, input *GetProductByIDInput) (*entity.Product, error) {
	product, err := pu.productRepository.GetByID(ctx, input.ID)
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	return product, nil
}
