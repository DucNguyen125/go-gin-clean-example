package main

import (
	"log"

	"base-gin-golang/cmd/wire"
	"base-gin-golang/config"
	"base-gin-golang/infra/postgresql"
	"base-gin-golang/middlewares"
	"base-gin-golang/pkg/logger"
	"base-gin-golang/routers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	cfg := loadEnvironment()
	if cfg.DebugMode {
		gin.SetMode("debug")
	} else {
		gin.SetMode("release")
	}
	// Init logger
	logger.Init(cfg)
	// Connect to database
	db, err := postgresql.ConnectPostgresql(cfg)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	if cfg.PostgreSQLMigrate {
		err = db.AutoMigrate()
		if err != nil {
			log.Fatal("Error migrating database")
		}
	}
	app, err := wire.InitApp(cfg, db)
	if err != nil {
		log.Fatal("Error create app")
	}
	// Middleware
	middleware := middlewares.NewMiddleware(
		app.JwtService,
		app.StringService,
		app.UserRepository,
	)
	vld := validator.New()
	cli := routers.InitRouter(
		cfg,
		middleware,
		&app,
		vld,
	)
	cli.Run()
}

func loadEnvironment() *config.Environment {
	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Fail loading environment variables: ", err)
	}
	return cfg
}
