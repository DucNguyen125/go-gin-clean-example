package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"base-gin-golang/domain/repository"
	"base-gin-golang/usecase/product"
)

func CreateProduct(productRepository repository.ProductRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input product.CreateProductInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := product.Create(productRepository, &input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusCreated, output)
	}
}

func GetProduct(productRepository repository.ProductRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input product.GetProductByIdInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input.Id = id
		output, err := product.GetById(productRepository, &input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListProduct(productRepository repository.ProductRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		output, err := product.GetList(productRepository, &input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateProduct(productRepository repository.ProductRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		input.Id = id
		output, err := product.Update(productRepository, &input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeleteProduct(productRepository repository.ProductRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input product.DeleteProductInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input.Id = id
		output, err := product.Delete(productRepository, &input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
