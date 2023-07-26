package postgresql

import (
	"base-gin-golang/config"
	"fmt"
	"math"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func ConnectPostgresql(cfg *config.Environment) (*Database, error) {
	var db *gorm.DB
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PostgresqlHost,
		cfg.PostgresqlUserName,
		cfg.PostgresqlPassword,
		cfg.PostgresqlDatabase,
		cfg.PostgresqlPort,
	)
	for i := 0; i < 3; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger(cfg),
		})
		if err != nil {
			fmt.Printf("attempt connecting the database...(%d)\n", i+1)
			// Retry connecting DB
			time.Sleep(time.Second * time.Duration(math.Pow(3, float64(i))))
		}
	}
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error dsn: %q", dsn))
	}
	return &Database{db}, nil
}

func (d Database) AutoMigrate() error {
	if err := initDatabase(d.DB); err != nil {
		return err
	}
	return nil
}
