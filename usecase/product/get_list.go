package product

import (
	"context"

	"base-gin-golang/domain/entity"
	"base-gin-golang/pkg/logger"
)

type GetListProductInput struct {
	PageIndex int    `query:"pageIndex"`
	PageSize  int    `query:"pageSize"`
	Order     string `query:"order"`
}

type GetListProductOutput struct {
	Body []*entity.Product
}

func (u *productUseCase) GetList(
	ctx context.Context,
	input *GetListProductInput,
) (*GetListProductOutput, error) {
	products, err := u.productRepository.GetList(ctx, entity.GetListProductOption{
		GetListOption: entity.GetListOption{
			PageIndex: input.PageIndex,
			PageSize:  input.PageSize,
			Order:     input.Order,
		},
	})
	if err != nil {
		logger.LogHandler(ctx, err)
		return nil, err
	}
	return &GetListProductOutput{Body: products}, nil
}
