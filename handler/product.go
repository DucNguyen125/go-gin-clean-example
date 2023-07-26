package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"base-gin-golang/usecase/product"
)

func CreateProduct(c *gin.Context, productUseCase product.UseCase) {
	var input product.CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	output, err := productUseCase.Create(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, output)
}

func GetProduct(c *gin.Context, productUseCase product.UseCase) {
	var input product.GetProductByIDInput
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = id
	output, err := productUseCase.GetByID(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, output)
}

func GetListProduct(c *gin.Context, productUseCase product.UseCase) {
	var input product.GetListProductInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
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
	output, err := productUseCase.GetList(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, output)
}

func UpdateProduct(c *gin.Context, productUseCase product.UseCase) {
	var input product.UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = id
	output, err := productUseCase.Update(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, output)
}

func DeleteProduct(c *gin.Context, productUseCase product.UseCase) {
	var input product.DeleteProductInput
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = id
	output, err := productUseCase.Delete(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}
