package v1

import (
	"base-gin-golang/handler"
	"base-gin-golang/usecase/product"

	"github.com/gin-gonic/gin"
)

func initProductRouter(
	r gin.IRouter,
	productUseCase product.UseCase,
) {
	r.POST("", func(ctx *gin.Context) {
		handler.CreateProduct(ctx, productUseCase)
	})
	r.GET("", func(ctx *gin.Context) {
		handler.GetListProduct(ctx, productUseCase)
	})
	r.GET("/:id", func(ctx *gin.Context) {
		handler.GetProduct(ctx, productUseCase)
	})
	r.PUT("/:id", func(ctx *gin.Context) {
		handler.UpdateProduct(ctx, productUseCase)
	})
	r.DELETE("/:id", func(ctx *gin.Context) {
		handler.DeleteProduct(ctx, productUseCase)
	})
}
