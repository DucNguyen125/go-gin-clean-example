package product

import (
	"base-gin-golang/domain/entity"
	"base-gin-golang/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UpdateProductInput struct {
	ID          int64
	ProductCode string `json:"productCode" binding:"required"`
	ProductName string `json:"productName" binding:"required"`
	Price       int    `json:"price" binding:"required"`
}

func (pu *productUseCase) Update(ctx *gin.Context, input *UpdateProductInput) (*entity.Product, error) {
	data := &entity.Product{}
	err := copier.Copy(data, input)
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	newProduct, err := pu.productRepository.Update(ctx, input.ID, data)
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	return newProduct, nil
}
