package routers

import (
	"base-gin-golang/domain/repository"
	"base-gin-golang/handler"

	"github.com/gin-gonic/gin"
)

func InitProductRouter(
	r gin.IRouter,
	productRepository repository.ProductRepository,
) {
	r.POST("/", handler.CreateProduct(productRepository))
	r.GET("/", handler.GetListProduct(productRepository))
	r.GET("/:id", handler.GetProduct(productRepository))
	r.PUT("/:id", handler.UpdateProduct(productRepository))
	r.DELETE("/:id", handler.DeleteProduct(productRepository))
}
