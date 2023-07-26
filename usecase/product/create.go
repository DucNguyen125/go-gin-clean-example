package product

import (
	"base-gin-golang/domain/entity"

	"base-gin-golang/domain/repository"

	"github.com/jinzhu/copier"
)

type CreateProductInput struct {
	ProductCode string `json:"productCode" binding:"required"`
	ProductName string `json:"productName" binding:"required"`
	Price       int    `json:"price" binding:"required"`
}

func Create(productRepository repository.ProductRepository, input *CreateProductInput) (*entity.Product, error) {
	data := &entity.Product{}
	err := copier.Copy(data, input)
	if err != nil {
		return nil, err
	}
	newProduct, err := productRepository.Create(data)
	if err != nil {
		return nil, err
	}
	return newProduct, nil
}
