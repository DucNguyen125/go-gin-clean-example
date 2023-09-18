package product

import (
	"base-gin-golang/domain/entity"
	"base-gin-golang/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (pu *productUseCase) CreateWithTransaction(ctx *gin.Context, input *CreateProductInput) (*entity.Product, error) {
	data := &entity.Product{}
	err := pu.dataService.Copy(data, input)
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	var newProduct *entity.Product
	err = pu.database.Transaction(func(tx *gorm.DB) error {
		ctx.Set("tx", tx)
		newProduct, err = pu.productRepository.Create(ctx, data)
		return err
	})
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	return newProduct, nil
}
