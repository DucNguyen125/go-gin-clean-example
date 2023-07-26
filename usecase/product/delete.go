package product

type DeleteProductInput struct {
	ID int64
}

type DeleteProductOutPut struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func (pu *productUseCase) Delete(input *DeleteProductInput) (*DeleteProductOutPut, error) {
	rowsAffected, err := pu.productRepository.Delete(input.ID)
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
