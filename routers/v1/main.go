package v1

import (
	"base-gin-golang/usecase/auth"
	"base-gin-golang/usecase/product"

	"github.com/gin-gonic/gin"
)

func InitV1Router(
	r *gin.RouterGroup,
	productUseCase product.UseCase,
	authUseCase auth.UseCase,
) {
	r.Use()
	{
		initProductRouter(r.Group("/products"), productUseCase)
		initAuthRouter(r.Group("/auth"), authUseCase)
	}
}
