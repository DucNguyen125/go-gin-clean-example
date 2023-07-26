package postgresql

import (
	"base-gin-golang/infra/postgresql/model"

	"gorm.io/gorm"
)

func initDatabase(db *gorm.DB) error {
	var err error
	if err = db.AutoMigrate(&model.Product{}); err != nil {
		return err
	}
	return nil
}
