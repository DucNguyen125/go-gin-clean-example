package auth

import (
	"context"

	"base-gin-golang/config"
	"base-gin-golang/domain/repository"
	jwtPkg "base-gin-golang/pkg/jwt"
	passwordPkg "base-gin-golang/pkg/password"
)

type UseCase interface {
	Login(ctx context.Context, input *LoginInput) (*LoginOutput, error)
}

type authUseCase struct {
	cfg             *config.Environment
	jwtService      jwtPkg.Service
	passwordService passwordPkg.Service
	userRepository  repository.UserRepository
}

func NewAuthUseCase(
	cfg *config.Environment,
	jwtService jwtPkg.Service,
	passwordService passwordPkg.Service,
	userRepository repository.UserRepository,
) UseCase {
	return &authUseCase{
		cfg:             cfg,
		userRepository:  userRepository,
		jwtService:      jwtService,
		passwordService: passwordService,
	}
}
