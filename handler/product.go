package handler

import (
	errorPkg "base-gin-golang/pkg/errors"
	errors "base-gin-golang/pkg/errors/custom"
	"base-gin-golang/pkg/pagination"
	"base-gin-golang/usecase/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context, productUseCase product.UseCase, errorService errorPkg.Service) {
	var input product.CreateProductInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		errValidate := errors.NewValidateError(ctx, input, err)
		ctx.JSON(http.StatusBadRequest, errValidate)
		return
	}
	output, err := productUseCase.Create(ctx, &input)
	if err != nil {
		errConverted := errorService.ParseInternalServer(err)
		ctx.JSON(errConverted.GetHTTPCode(), errConverted)
		return
	}
	ctx.JSON(http.StatusCreated, output)
}

func GetProduct(ctx *gin.Context, productUseCase product.UseCase, errorService errorPkg.Service) {
	var input product.GetProductByIDInput
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		errValidate := errors.NewValidateError(ctx, input, err)
		ctx.JSON(http.StatusBadRequest, errValidate)
		return
	}
	input.ID = id
	output, err := productUseCase.GetByID(ctx, &input)
	if err != nil {
		errConverted := errorService.ParseInternalServer(err)
		ctx.JSON(errConverted.GetHTTPCode(), errConverted)
		return
	}
	ctx.JSON(http.StatusOK, output)
}

func GetListProduct(ctx *gin.Context, productUseCase product.UseCase, errorService errorPkg.Service) {
	var input product.GetListProductInput
	if err := ctx.ShouldBind(&input); err != nil {
		errValidate := errors.NewValidateError(ctx, input, err)
		ctx.JSON(http.StatusBadRequest, errValidate)
		return
	}
	pageIndex, pageSize, order := pagination.GetDefaultPagination(
		input.PageIndex, input.PageSize, input.Order,
	)
	input.PageIndex = pageIndex
	input.PageSize = pageSize
	input.Order = order
	output, err := productUseCase.GetList(ctx, &input)
	if err != nil {
		errConverted := errorService.ParseInternalServer(err)
		ctx.JSON(errConverted.GetHTTPCode(), errConverted)
		return
	}
	ctx.JSON(http.StatusOK, output)
}

func UpdateProduct(ctx *gin.Context, productUseCase product.UseCase, errorService errorPkg.Service) {
	var input product.UpdateProductInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		errValidate := errors.NewValidateError(ctx, input, err)
		ctx.JSON(http.StatusBadRequest, errValidate)
		return
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = id
	output, err := productUseCase.Update(ctx, &input)
	if err != nil {
		errConverted := errorService.ParseInternalServer(err)
		ctx.JSON(errConverted.GetHTTPCode(), errConverted)
		return
	}
	ctx.JSON(http.StatusOK, output)
}

func DeleteProduct(ctx *gin.Context, productUseCase product.UseCase, errorService errorPkg.Service) {
	var input product.DeleteProductInput
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		errValidate := errors.NewValidateError(ctx, input, err)
		ctx.JSON(http.StatusBadRequest, errValidate)
		return
	}
	input.ID = id
	output, err := productUseCase.Delete(ctx, &input)
	if err != nil {
		errConverted := errorService.ParseInternalServer(err)
		ctx.JSON(errConverted.GetHTTPCode(), errConverted)
		return
	}
	ctx.JSON(http.StatusOK, output)
}
