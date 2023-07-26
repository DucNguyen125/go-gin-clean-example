package postgresql

import (
	"base-gin-golang/infra/postgresql/model"

	"gorm.io/gorm"
)

func initDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.Product{}); err != nil {
		return err
	}
	err := db.AutoMigrate(&model.User{})
	return err
}
