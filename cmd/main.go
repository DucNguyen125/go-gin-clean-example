package main

import (
	"base-gin-golang/config"
	"base-gin-golang/infra/postgresql"
	"base-gin-golang/infra/postgresql/repository"
	dataPkg "base-gin-golang/pkg/data"
	"base-gin-golang/pkg/logger"
	"base-gin-golang/routers"
	"base-gin-golang/usecase/product"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type App struct {
	config   *config.Environment
	database *postgresql.Database
}

func main() {
	cfg := loadEnvironment()
	gin.SetMode(cfg.RunMode)
	// Init logger
	logger.Init(cfg)
	// Connect to database
	db, err := postgresql.ConnectPostgresql(cfg)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	app := &App{
		config:   cfg,
		database: db,
	}
	err = app.database.AutoMigrate()
	if err != nil {
		log.Fatal("Error migrating database")
	}
	// Service
	dataService := dataPkg.NewDataService()
	// Repository
	productRepository := repository.NewProductRepository(app.database, dataService)
	// UseCase
	productUseCase := product.NewProductUseCase(productRepository, dataService)
	router := routers.InitRouter(
		app.config,
		productUseCase,
	)
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.Port),
		ReadHeaderTimeout: 3 * time.Second, //nolint:gomnd // common
		Handler:           router,
	}
	log.Printf("[info] start http server listening: %d", app.config.Port)
	if err = server.ListenAndServe(); err != nil {
		log.Fatal("Fail to start error server")
	}
}

func loadEnvironment() *config.Environment {
	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Fail loading environment variables: ", err)
	}
	return cfg
}
