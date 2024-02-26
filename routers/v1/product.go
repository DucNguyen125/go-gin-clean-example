package v1

import (
	"context"

	"base-gin-golang/handler"
	routerPkg "base-gin-golang/pkg/router"
	"base-gin-golang/usecase/product"
)

func initProductRouter(rg *routerPkg.RouterGroup, hdl *handler.Handler) {
	routerPkg.POST("", rg,
		func(ctx context.Context, input *product.CreateProductInput) (*product.CreateProductOutput, error) {
			return hdl.CreateProduct(ctx, input)
		})
	routerPkg.GET("", rg,
		func(ctx context.Context, input *product.GetListProductInput) (*product.GetListProductOutput, error) {
			return hdl.GetListProduct(ctx, input)
		})
	routerPkg.GET("/{id}", rg,
		func(ctx context.Context, input *product.GetProductByIDInput) (*product.GetProductByIDOutput, error) {
			return hdl.GetProduct(ctx, input)
		})
	routerPkg.PUT("/{id}", rg,
		func(ctx context.Context, input *product.UpdateProductInput) (*product.UpdateProductOutput, error) {
			return hdl.UpdateProduct(ctx, input)
		})
	routerPkg.DELETE("/{id}", rg,
		func(ctx context.Context, input *product.DeleteProductInput) (*product.DeleteProductOutput, error) {
			return hdl.DeleteProduct(ctx, input)
		})
}
