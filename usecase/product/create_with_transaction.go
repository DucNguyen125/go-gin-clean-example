package product

import (
	"context"

	"base-gin-golang/constants"
	"base-gin-golang/domain/entity"
	"base-gin-golang/pkg/logger"

	"gorm.io/gorm"
)

func (u *productUseCase) CreateWithTransaction(
	ctx context.Context,
	input *CreateProductInput,
) (*CreateProductOutput, error) {
	data := &entity.Product{}
	err := u.dataService.Copy(data, &input.Body)
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	var newProduct *entity.Product
	err = u.database.Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, constants.ContextKeyTransaction, tx)
		newProduct, err = u.productRepository.Create(ctx, data)
		return err
	})
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	return &CreateProductOutput{Body: newProduct}, nil
}
