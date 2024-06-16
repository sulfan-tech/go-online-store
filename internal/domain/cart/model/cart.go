package model

import (
	"go-online-store/internal/domain/product/model"
	"time"
)

type Cart struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	CustomerID uint       `json:"customer_id" gorm:"not null"`
	Items      []CartItem `json:"items" gorm:"foreignKey:CartID"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type CartItem struct {
	ID        uint          `json:"id" gorm:"primaryKey"`
	CartID    uint          `json:"cart_id" gorm:"not null"`
	ProductID uint          `json:"product_id" gorm:"not null"`
	Quantity  uint          `json:"quantity" gorm:"not null"`
	Product   model.Product `json:"product" gorm:"foreignKey:ProductID"`
}

func (CartItem) TableName() string {
	return "CartItem"
}

func (Cart) TableName() string {
	return "Cart"
}
