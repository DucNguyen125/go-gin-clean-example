package v1

import (
	"base-gin-golang/handler"
	routerPkg "base-gin-golang/pkg/router"
)

func InitV1Router(r *routerPkg.RouterGroup, hdl *handler.Handler) {
	initAuthRouter(r.Group("/auth"), hdl)
	initProductRouter(r.Group("/products"), hdl)
}
