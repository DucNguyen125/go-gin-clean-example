package routers

import (
	"base-gin-golang/config"
	"base-gin-golang/domain/repository"
	"base-gin-golang/middlewares"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(
	config *config.Environment,
	productRepository repository.ProductRepository,
) *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: strings.Split(config.CorsAllowOrigins, ","),
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Access-Control-Allow-Headers",
			"Authorization",
			"X-XSRF-TOKEN",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, //nolint:gomnd // common
	}))
	router.Use(middlewares.RestLogger)
	router.Use(gin.Recovery())
	apiRouter := router.Group("/api")
	apiRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	apiRouter.Use()
	{
		InitProductRouter(apiRouter.Group("/products"), productRepository)
	}
	return router
}
