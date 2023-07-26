package wire

import (
	"base-gin-golang/domain/repository"
	dataPkg "base-gin-golang/pkg/data"
	errorPkg "base-gin-golang/pkg/errors"
	jwtPkg "base-gin-golang/pkg/jwt"
	passwordPkg "base-gin-golang/pkg/password"
	stringPkg "base-gin-golang/pkg/string"
	"base-gin-golang/usecase/auth"
	"base-gin-golang/usecase/product"
)

type App struct {
	// Service
	DataService     dataPkg.Service
	StringService   stringPkg.Service
	JwtService      jwtPkg.Service
	PasswordService passwordPkg.Service
	ErrorService    errorPkg.Service
	// Repository
	ProductRepository repository.ProductRepository
	UserRepository    repository.UserRepository
	// UseCase
	ProductUseCase product.UseCase
	AuthUseCase    auth.UseCase
}

func newApp(
	// Service
	dataService dataPkg.Service,
	stringService stringPkg.Service,
	jwtService jwtPkg.Service,
	passwordService passwordPkg.Service,
	errorService errorPkg.Service,
	// Repository
	productRepository repository.ProductRepository,
	userRepository repository.UserRepository,
	// UseCase
	productUseCase product.UseCase,
	authUseCase auth.UseCase,
) App {
	return App{
		// Service
		DataService:     dataService,
		StringService:   stringService,
		JwtService:      jwtService,
		PasswordService: passwordService,
		ErrorService:    errorService,
		// Repository
		ProductRepository: productRepository,
		UserRepository:    userRepository,
		// UseCase
		ProductUseCase: productUseCase,
		AuthUseCase:    authUseCase,
	}
}
