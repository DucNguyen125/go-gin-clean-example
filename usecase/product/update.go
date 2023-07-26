package product

import (
	"base-gin-golang/domain/entity"

	"base-gin-golang/domain/repository"

	"github.com/jinzhu/copier"
)

type UpdateProductInput struct {
	Id          int64
	ProductCode string `json:"productCode" binding:"required"`
	ProductName string `json:"productName" binding:"required"`
	Price       int    `json:"price" binding:"required"`
}

func Update(productRepository repository.ProductRepository, input *UpdateProductInput) (*entity.Product, error) {
	data := &entity.Product{}
	err := copier.Copy(data, input)
	if err != nil {
		return nil, err
	}
	newProduct, err := productRepository.Update(input.Id, data)
	if err != nil {
		return nil, err
	}
	return newProduct, nil
}
