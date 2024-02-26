package router

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func handle[I, O any](
	httpMethod string,
	relativePath string,
	group *RouterGroup,
	handleFunc func(context.Context, *I) (*O, error),
) {
	absolutePath := group.calculateAbsolutePath(relativePath)
	huma.Register(group.routerService.HumaAPI, huma.Operation{
		Method:           httpMethod,
		Path:             absolutePath,
		Hidden:           !group.routerService.Config.DebugMode,
		SkipValidateBody: true,
	}, handleFunc)
}

func POST[I, O any](
	relativePath string,
	group *RouterGroup,
	handlerFunc func(context.Context, *I) (*O, error),
) {
	handle(http.MethodPost, relativePath, group, handlerFunc)
}

func GET[I, O any](
	relativePath string,
	group *RouterGroup,
	handlerFunc func(context.Context, *I) (*O, error),
) {
	handle(http.MethodGet, relativePath, group, handlerFunc)
}

func DELETE[I, O any](
	relativePath string,
	group *RouterGroup,
	handlerFunc func(context.Context, *I) (*O, error),
) {
	handle(http.MethodDelete, relativePath, group, handlerFunc)
}

func PATCH[I, O any](
	relativePath string,
	group *RouterGroup,
	handlerFunc func(context.Context, *I) (*O, error),
) {
	handle(http.MethodPatch, relativePath, group, handlerFunc)
}

func PUT[I, O any](
	relativePath string,
	group *RouterGroup,
	handlerFunc func(context.Context, *I) (*O, error),
) {
	handle(http.MethodPut, relativePath, group, handlerFunc)
}
