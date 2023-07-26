package internal

import "time"

type User struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
