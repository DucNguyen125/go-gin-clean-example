package main

import (
	"base-gin-golang/config"
	"base-gin-golang/infra/postgresql"
	"base-gin-golang/infra/postgresql/repository"
	dataPkg "base-gin-golang/pkg/data"
	"base-gin-golang/pkg/logger"
	"base-gin-golang/routers"
	"base-gin-golang/usecase/product"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type App struct {
	config   *config.Environment
	database *postgresql.Database
}

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
	productUseCase := product.NewProductUseCase(productRepository, dataService, app.database)
	router := routers.InitRouter(
		app.config,
		productUseCase,
	)
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.Port),
		ReadHeaderTimeout: 3 * time.Second, //nolint:gomnd // common
		Handler:           router,
	}
	done := make(chan bool)
	go func() {
		if subErr := gracefulShutDown(app.config, done, server); subErr != nil {
			logrus.Errorf("Stop server shutdown error: %v", err.Error())
			return
		}
		logrus.Info("Stopped serving on Services")
	}()
	log.Printf("Start HTTP Server, Listening: %d", app.config.Port)
	if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Start HTTP Server Failed. Error: %s", err.Error())
	}
	<-done
	logrus.Info("Stopped backend application.")
}

func loadEnvironment() *config.Environment {
	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Fail loading environment variables: ", err)
	}
	return cfg
}

func gracefulShutDown(config *config.Environment, quit chan bool, server *http.Server) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.SystemShutdownTimeOutSecond)*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		return err
	}
	close(quit)
	return nil
}
