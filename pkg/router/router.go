package router

import (
	"base-gin-golang/config"

	"github.com/danielgtaylor/huma/v2"
)

type RouterService struct {
	Config  *config.Environment
	HumaAPI huma.API
}

func NewRouterService(cfg *config.Environment, humaApi huma.API) *RouterService {
	return &RouterService{
		Config:  cfg,
		HumaAPI: humaApi,
	}
}

func (r *RouterService) Group(relativePath string) *RouterGroup {
	return &RouterGroup{
		routerService: r,
		basePath:      relativePath,
	}
}
