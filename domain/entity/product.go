package entity

import "time"

type Product struct {
	ID          int       `json:"id"`
	ProductCode string    `json:"productCode"`
	ProductName string    `json:"productName"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type GetListProductOption struct {
	GetListOption
}
