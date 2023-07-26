package model

import (
	"base-gin-golang/infra/postgresql/model/internal"

	"gorm.io/gorm"
)

type BaseModel struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

type Product struct {
	internal.Product
	BaseModel
}

type User struct {
	internal.User
	BaseModel
}
