package v1

import (
	"context"

	"base-gin-golang/handler"
	routerPkg "base-gin-golang/pkg/router"
	"base-gin-golang/usecase/auth"
)

func initAuthRouter(rg *routerPkg.RouterGroup, hdl *handler.Handler) {
	routerPkg.POST("/login", rg,
		func(ctx context.Context, input *auth.LoginInput) (*auth.LoginOutput, error) {
			return hdl.Login(ctx, input)
		})
}
