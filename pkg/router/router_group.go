package router

import "path"

type RouterGroup struct {
	routerService *RouterService
	basePath      string
}

func NewRouterGroup(routerService *RouterService, basePath string) *RouterGroup {
	return &RouterGroup{
		routerService: routerService,
		basePath:      basePath,
	}
}

func (group *RouterGroup) Group(relativePath string) *RouterGroup {
	return &RouterGroup{
		routerService: group.routerService,
		basePath:      group.calculateAbsolutePath(relativePath),
	}
}

func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}

func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	if lastChar(relativePath) == '/' && lastChar(finalPath) != '/' {
		return finalPath + "/"
	}
	return finalPath
}

func (group *RouterGroup) calculateAbsolutePath(relativePath string) string {
	return joinPaths(group.basePath, relativePath)
}
