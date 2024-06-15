package model

import (
	"time"
)

type CartItem struct {
	CartItemID uint `gorm:"primaryKey;column:cart_item_id"`
	CartID     uint `gorm:"index;column:cart_id"`
	ProductID  uint `gorm:"index;column:product_id"`
	Quantity   uint
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

type Cart struct {
	CartID     uint       `gorm:"primaryKey;column:cart_id"`
	CustomerID uint       `gorm:"index;column:customer_id"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  time.Time  `gorm:"column:updated_at"`
	Items      []CartItem `gorm:"foreignKey:CartID"`
}

func (CartItem) TableName() string {
	return "CartItem"
}

func (Cart) TableName() string {
	return "Cart"
}
