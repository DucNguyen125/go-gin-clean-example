package v1

import (
	"base-gin-golang/handler"
	"base-gin-golang/usecase/auth"

	"github.com/gin-gonic/gin"
)

func initAuthRouter(
	r gin.IRouter,
	authUseCase auth.UseCase,
) {
	r.POST("/login", func(context *gin.Context) {
		handler.Login(context, authUseCase)
	})
}
