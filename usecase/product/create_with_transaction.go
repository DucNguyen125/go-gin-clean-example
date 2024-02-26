package product

import (
	"context"

	"base-gin-golang/config"
	"base-gin-golang/domain/entity"
	"base-gin-golang/pkg/logger"

	"gorm.io/gorm"
)

func (pu *productUseCase) CreateWithTransaction(
	ctx context.Context,
	input *CreateProductInput,
) (*CreateProductOutput, error) {
	data := &entity.Product{}
	err := pu.dataService.Copy(data, &input.Body)
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	var newProduct *entity.Product
	err = pu.database.Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, config.ContextKeyTransaction, tx)
		newProduct, err = pu.productRepository.Create(ctx, data)
		return err
	})
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	return &CreateProductOutput{Body: newProduct}, nil
}
