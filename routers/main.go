package routers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"base-gin-golang/cmd/wire"
	"base-gin-golang/config"
	"base-gin-golang/handler"
	"base-gin-golang/middlewares"
	routerPkg "base-gin-golang/pkg/router"
	v1Routers "base-gin-golang/routers/v1"
	"base-gin-golang/validations"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

// Options for the CLI.
type Options struct{}

func InitRouter(
	cfg *config.Environment,
	middleware middlewares.Middleware,
	app *wire.App,
	vld *validator.Validate,
) huma.CLI {
	humaCli := huma.NewCLI(func(hooks huma.Hooks, option *Options) {
		ginEngine := gin.New()
		ginEngine.Use(cors.New(cors.Config{
			AllowOrigins: strings.Split(cfg.CorsAllowOrigins, ","),
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
		ginEngine.Use(
			middleware.RestLogger,
			gin.Recovery(),
			middleware.Authentication,
		)
		// Validations
		vld.RegisterValidation("customEmail", validations.CustomEmail) //nolint:errcheck // no-error

		humaApi := humagin.New(ginEngine, huma.DefaultConfig("Base Gin Golang", "2.0.0"))
		routerService := routerPkg.NewRouterService(cfg, humaApi)
		apiRouterGroup := routerService.Group("/api")
		type (
			Input struct{}
			Body  struct {
				Message string `json:"message"`
			}
			Output struct {
				Body Body
			}
		)
		routerPkg.GET("/ping", apiRouterGroup,
			func(ctx context.Context, input *Input) (*Output, error) {
				return &Output{
					Body: Body{
						Message: "pong",
					},
				}, nil
			})
		hdl := handler.NewHandler(cfg, app, vld)
		v1Routers.InitV1Router(apiRouterGroup.Group("/v1"), hdl)
		server := &http.Server{
			Addr:              fmt.Sprintf(":%d", cfg.Port),
			ReadHeaderTimeout: 3 * time.Second, //nolint:gomnd // common
			Handler:           ginEngine,
		}
		hooks.OnStart(func() {
			logrus.Infof("Start HTTP Server, Listening: %d", cfg.Port)
			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				logrus.Fatalf("Start HTTP Server Failed. Error: %s", err.Error())
			}
		})
		hooks.OnStop(func() {
			ctx, cancel := context.WithTimeout(
				context.Background(),
				time.Duration(cfg.SystemShutdownTimeOutSecond)*time.Second,
			)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				logrus.Errorf("Stop server shutdown error: %v", err.Error())
			}
			logrus.Info("Stopped backend application.")
		})
	})
	return humaCli
}
