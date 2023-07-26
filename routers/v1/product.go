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
	r.POST("/", func(context *gin.Context) {
		handler.CreateProduct(context, productUseCase)
	})
	r.GET("/", func(context *gin.Context) {
		handler.GetListProduct(context, productUseCase)
	})
	r.GET("/:id", func(context *gin.Context) {
		handler.GetProduct(context, productUseCase)
	})
	r.PUT("/:id", func(context *gin.Context) {
		handler.UpdateProduct(context, productUseCase)
	})
	r.DELETE("/:id", func(context *gin.Context) {
		handler.DeleteProduct(context, productUseCase)
	})
}
