package v1

import (
	errorPkg "base-gin-golang/pkg/errors"
	"base-gin-golang/usecase/auth"
	"base-gin-golang/usecase/product"

	"github.com/gin-gonic/gin"
)

func InitV1Router(
	r *gin.RouterGroup,
	productUseCase product.UseCase,
	authUseCase auth.UseCase,
	errorService errorPkg.Service,
) {
	r.Use()
	{
		initProductRouter(r.Group("/products"), productUseCase, errorService)
		initAuthRouter(r.Group("/auth"), authUseCase, errorService)
	}
}
