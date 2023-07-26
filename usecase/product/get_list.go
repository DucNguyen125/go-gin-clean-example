package product

import (
	"base-gin-golang/domain/entity"
)

type GetListProductInput struct {
	PageIndex int     `form:"pageIndex"`
	PageSize  int     `form:"pageSize"`
	Order     *string `form:"order"`
}

func (pu *productUseCase) GetList(input *GetListProductInput) ([]*entity.Product, error) {
	products, err := pu.productRepository.GetList(entity.GetListProductOption{
		GetListOption: entity.GetListOption{
			PageIndex: input.PageIndex,
			PageSize:  input.PageSize,
			Order:     input.Order,
		},
	})
	if err != nil {
		return nil, err
	}
	return products, nil
}
