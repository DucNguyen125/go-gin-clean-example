package product

import (
	"context"

	"base-gin-golang/domain/repository"
	"base-gin-golang/infra/postgresql"
	dataPkg "base-gin-golang/pkg/data"
)

type UseCase interface {
	Create(ctx context.Context, input *CreateProductInput) (*CreateProductOutput, error)
	Delete(ctx context.Context, input *DeleteProductInput) (*DeleteProductOutput, error)
	GetByID(ctx context.Context, input *GetProductByIDInput) (*GetProductByIDOutput, error)
	GetList(ctx context.Context, input *GetListProductInput) (*GetListProductOutput, error)
	Update(ctx context.Context, input *UpdateProductInput) (*UpdateProductOutput, error)
	CreateWithTransaction(ctx context.Context, input *CreateProductInput) (*CreateProductOutput, error)
}

type productUseCase struct {
	productRepository repository.ProductRepository
	dataService       dataPkg.Service
	database          *postgresql.Database
}

func NewProductUseCase(
	productRepository repository.ProductRepository,
	dataService dataPkg.Service,
	database *postgresql.Database,
) UseCase {
	return &productUseCase{
		productRepository: productRepository,
		dataService:       dataService,
		database:          database,
	}
}
