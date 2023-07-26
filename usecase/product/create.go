package product

import (
	"base-gin-golang/domain/entity"

	"github.com/gin-gonic/gin"
)

type CreateProductInput struct {
	ProductCode string `json:"productCode" binding:"required"`
	ProductName string `json:"productName" binding:"required"`
	Price       int    `json:"price" binding:"required"`
}

func (pu *productUseCase) Create(ctx *gin.Context, input *CreateProductInput) (*entity.Product, error) {
	data := &entity.Product{}
	err := pu.dataService.Copy(data, input)
	if err != nil {
		return nil, err
	}
	newProduct, err := pu.productRepository.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return newProduct, nil
}
