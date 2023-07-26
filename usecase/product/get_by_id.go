package product

import (
	"base-gin-golang/domain/entity"

	"github.com/gin-gonic/gin"
)

type GetProductByIDInput struct {
	ID int64
}

func (pu *productUseCase) GetByID(ctx *gin.Context, input *GetProductByIDInput) (*entity.Product, error) {
	product, err := pu.productRepository.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}
