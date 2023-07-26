package product

import (
	"base-gin-golang/domain/entity"

	"github.com/gin-gonic/gin"
)

type GetListProductInput struct {
	PageIndex int     `form:"pageIndex"`
	PageSize  int     `form:"pageSize"`
	Order     *string `form:"order"`
}

func (pu *productUseCase) GetList(ctx *gin.Context, input *GetListProductInput) ([]*entity.Product, error) {
	products, err := pu.productRepository.GetList(ctx, entity.GetListProductOption{
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
