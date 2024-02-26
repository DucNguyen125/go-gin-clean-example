package handler

import (
	"context"

	errors "base-gin-golang/pkg/errors/custom"
	"base-gin-golang/pkg/pagination"
	"base-gin-golang/usecase/product"
)

func (h *Handler) CreateProduct(
	ctx context.Context,
	input *product.CreateProductInput,
) (*product.CreateProductOutput, error) {
	if err := h.Validator.Struct(&input.Body); err != nil {
		errValidate := errors.NewValidateError(ctx, input, err)
		return nil, errValidate
	}
	output, err := h.App.ProductUseCase.Create(ctx, input)
	if err != nil {
		errConverted := h.App.ErrorService.ParseInternalServer(err)
		return nil, errConverted
	}
	return output, err
}

func (h *Handler) GetProduct(
	ctx context.Context,
	input *product.GetProductByIDInput,
) (*product.GetProductByIDOutput, error) {
	output, err := h.App.ProductUseCase.GetByID(ctx, input)
	if err != nil {
		errConverted := h.App.ErrorService.ParseInternalServer(err)
		return nil, errConverted
	}
	return output, nil
}

func (h *Handler) GetListProduct(
	ctx context.Context,
	input *product.GetListProductInput,
) (*product.GetListProductOutput, error) {
	pageIndex, pageSize, order := pagination.GetDefaultPagination(
		input.PageIndex, input.PageSize, input.Order,
	)
	input.PageIndex = pageIndex
	input.PageSize = pageSize
	input.Order = order
	output, err := h.App.ProductUseCase.GetList(ctx, input)
	if err != nil {
		errConverted := h.App.ErrorService.ParseInternalServer(err)
		return nil, errConverted
	}
	return output, nil
}

func (h *Handler) UpdateProduct(
	ctx context.Context,
	input *product.UpdateProductInput,
) (*product.UpdateProductOutput, error) {
	if err := h.Validator.Struct(&input.Body); err != nil {
		errValidate := errors.NewValidateError(ctx, input, err)
		return nil, errValidate
	}
	output, err := h.App.ProductUseCase.Update(ctx, input)
	if err != nil {
		errConverted := h.App.ErrorService.ParseInternalServer(err)
		return nil, errConverted
	}
	return output, nil
}

func (h *Handler) DeleteProduct(
	ctx context.Context,
	input *product.DeleteProductInput,
) (*product.DeleteProductOutput, error) {
	output, err := h.App.ProductUseCase.Delete(ctx, input)
	if err != nil {
		errConverted := h.App.ErrorService.ParseInternalServer(err)
		return nil, errConverted
	}
	return output, nil
}
