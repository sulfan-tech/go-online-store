package model

import (
	"time"
)

type Product struct {
	ProductId string    `json:"product_id" gorm:"column:product_id;not null"`
	Name      string    `json:"name" gorm:"column:name;not null"`
	Category  string    `json:"category" gorm:"column:category;not null"`
	Price     float64   `json:"price" gorm:"column:price;not null"`
	Stok      int       `json:"stok" gorm:"column:stok;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Product) TableName() string {
	return "Product"
}
