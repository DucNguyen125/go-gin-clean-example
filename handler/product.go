package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"base-gin-golang/usecase/product"
)

func CreateProduct(ctx *gin.Context, productUseCase product.UseCase) {
	var input product.CreateProductInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	output, err := productUseCase.Create(ctx, &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, output)
}

func GetProduct(ctx *gin.Context, productUseCase product.UseCase) {
	var input product.GetProductByIDInput
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = id
	output, err := productUseCase.GetByID(ctx, &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, output)
}

func GetListProduct(ctx *gin.Context, productUseCase product.UseCase) {
	var input product.GetListProductInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if input.PageIndex == 0 {
		input.PageIndex = 1
	}
	if input.PageSize == 0 {
		input.PageSize = 20
	}
	if input.Order == nil || *input.Order == "" {
		defaultOrder := "id ASC"
		input.Order = &defaultOrder
	}
	output, err := productUseCase.GetList(ctx, &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, output)
}

func UpdateProduct(ctx *gin.Context, productUseCase product.UseCase) {
	var input product.UpdateProductInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
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
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, output)
}

func DeleteProduct(ctx *gin.Context, productUseCase product.UseCase) {
	var input product.DeleteProductInput
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = id
	output, err := productUseCase.Delete(ctx, &input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, output)
}
