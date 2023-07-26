package routers

import (
	"base-gin-golang/config"
	"base-gin-golang/middlewares"
	"base-gin-golang/usecase/product"
	"net/http"
	"strings"
	"time"

	v1Routers "base-gin-golang/routers/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(
	config *config.Environment,
	productUseCase product.UseCase,
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
		ExposeHeaders: []string{
			"Content-Disposition",
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
	v1Routers.InitV1Router(
		apiRouter.Group("/v1"),
		productUseCase,
	)
	return router
}
