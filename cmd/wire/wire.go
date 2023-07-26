//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package wire

import (
	"base-gin-golang/config"
	"base-gin-golang/infra/postgresql"
	"base-gin-golang/infra/postgresql/repository"
	dataPkg "base-gin-golang/pkg/data"
	errorPkg "base-gin-golang/pkg/errors"
	jwtPkg "base-gin-golang/pkg/jwt"
	passwordPkg "base-gin-golang/pkg/password"
	stringPkg "base-gin-golang/pkg/string"
	"base-gin-golang/usecase/auth"
	"base-gin-golang/usecase/product"

	"github.com/google/wire"
)

func InitApp(config *config.Environment, database *postgresql.Database) (App, error) {
	panic(wire.Build(
		// Service
		wire.NewSet(dataPkg.NewDataService),
		wire.NewSet(stringPkg.NewStringService),
		wire.NewSet(jwtPkg.NewJwtService),
		wire.NewSet(passwordPkg.NewPasswordService),
		wire.NewSet(errorPkg.NewErrorService),
		// Repository
		wire.NewSet(repository.NewProductRepository),
		wire.NewSet(repository.NewUserRepository),
		// UseCase
		wire.NewSet(product.NewProductUseCase),
		wire.NewSet(auth.NewAuthUseCase),
		newApp,
	))
}
