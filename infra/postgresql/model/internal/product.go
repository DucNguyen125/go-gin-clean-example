package internal

import "time"

type Product struct {
	ID          int       `gorm:"column:id;primaryKey;autoIncrement"`
	ProductCode string    `gorm:"column:product_code"`
	ProductName string    `gorm:"column:product_name"`
	Price       int       `gorm:"column:price"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
