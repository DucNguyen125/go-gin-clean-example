package auth

import (
	"base-gin-golang/config"
	"base-gin-golang/domain/repository"
	jwtPkg "base-gin-golang/pkg/jwt"
	stringPkg "base-gin-golang/pkg/string"

	"github.com/gin-gonic/gin"
)

type UseCase interface {
	Login(ctx *gin.Context, input *LoginInput) (*LoginOutput, error)
}

type authUseCase struct {
	cfg            config.Environment
	jwtService     jwtPkg.Service
	stringService  stringPkg.Service
	userRepository repository.UserRepository
}

func NewAuthUseCase(
	cfg config.Environment,
	jwtService jwtPkg.Service,
	stringService stringPkg.Service,
	userRepository repository.UserRepository,
) UseCase {
	return &authUseCase{
		cfg:            cfg,
		userRepository: userRepository,
		jwtService:     jwtService,
		stringService:  stringService,
	}
}
