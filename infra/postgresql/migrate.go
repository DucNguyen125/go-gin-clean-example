package postgresql

import (
	"base-gin-golang/infra/postgresql/model"

	"gorm.io/gorm"
)

func initDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Product{})
	return err
}
