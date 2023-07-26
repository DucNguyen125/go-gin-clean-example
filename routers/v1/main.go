package v1

import (
	"base-gin-golang/usecase/product"

	"github.com/gin-gonic/gin"
)

func InitV1Router(
	r *gin.RouterGroup,
	productUseCase product.UseCase,
) {
	r.Use()
	{
		initProductRouter(r.Group("/products"), productUseCase)
	}
}
