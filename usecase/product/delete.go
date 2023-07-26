package product

import (
	"base-gin-golang/domain/repository"
)

type DeleteProductInput struct {
	ID int64
}

type DeleteProductOutPut struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func Delete(productRepository repository.ProductRepository, input *DeleteProductInput) (*DeleteProductOutPut, error) {
	rowsAffected, err := productRepository.Delete(input.ID)
	if err != nil {
		return &DeleteProductOutPut{
			RowsAffected: 0,
		}, err
	}
	output := &DeleteProductOutPut{
		RowsAffected: rowsAffected,
	}
	return output, nil
}
