package middlewares

import (
	"base-gin-golang/domain/repository"
	jwtPkg "base-gin-golang/pkg/jwt"
	stringPkg "base-gin-golang/pkg/string"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	RestLogger(ctx *gin.Context)
	Authentication(ctx *gin.Context)
}

type middleware struct {
	jwtService     jwtPkg.Service
	stringService  stringPkg.Service
	userRepository repository.UserRepository
}

func NewMiddleware(
	jwtService jwtPkg.Service,
	stringService stringPkg.Service,
	userRepository repository.UserRepository,
) Middleware {
	return &middleware{
		jwtService:     jwtService,
		stringService:  stringService,
		userRepository: userRepository,
	}
}
