package product

import (
	"base-gin-golang/domain/entity"
	"base-gin-golang/domain/repository"
	dataPkg "base-gin-golang/pkg/data"
)

type UseCase interface {
	Create(input *CreateProductInput) (*entity.Product, error)
	Delete(input *DeleteProductInput) (*DeleteProductOutPut, error)
	GetByID(input *GetProductByIDInput) (*entity.Product, error)
	GetList(input *GetListProductInput) ([]*entity.Product, error)
	Update(input *UpdateProductInput) (*entity.Product, error)
}

type productUseCase struct {
	productRepository repository.ProductRepository
	dataService       dataPkg.Service
}

func NewProductUseCase(
	productRepository repository.ProductRepository,
	dataService dataPkg.Service,
) UseCase {
	return &productUseCase{
		productRepository: productRepository,
		dataService:       dataService,
	}
}
