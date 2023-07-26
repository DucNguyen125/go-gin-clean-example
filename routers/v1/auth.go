package v1

import (
	"base-gin-golang/handler"
	errorPkg "base-gin-golang/pkg/errors"
	"base-gin-golang/usecase/auth"

	"github.com/gin-gonic/gin"
)

func initAuthRouter(
	r gin.IRouter,
	authUseCase auth.UseCase,
	errorService errorPkg.Service,
) {
	r.POST("/login", func(context *gin.Context) {
		handler.Login(context, authUseCase, errorService)
	})
}
